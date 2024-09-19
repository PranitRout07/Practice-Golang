package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type ApiError struct {
	Error string `json:"error"`
}
type apiFunc func(w http.ResponseWriter, r *http.Request) error

func makeHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Println(r.Response)
			
			WriteJSON(w,http.StatusNotFound,ApiError{Error: err.Error()})
			// e := ApiError{Error: err.Error()}
			// WriteJSON(w, r.Response.StatusCode, e)
		}
	}

}

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listerAddr string,store Storage) *APIServer {
	return &APIServer{listenAddr: listerAddr,store: store}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	log.Println("Running on port:", s.listenAddr)
	router.HandleFunc("/account", makeHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}",withJWTAuth(makeHandleFunc(s.handleAccountByID)))
	router.HandleFunc("/transfer",makeHandleFunc(s.handleTransfer))
	router.HandleFunc("/login",makeHandleFunc(s.handleLogin))
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccounts(w, r)
	} else if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	} else if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("method not allowed::%s", r.Method)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	log.Println("Create called")
	data := r.Body
	byteData, err := io.ReadAll(data)
	if err != nil {
		log.Println("ERROR::", err)
	}
	var user AccountReq
	err = json.Unmarshal(byteData, &user)
	if err != nil {
		log.Println("ERROR::", err)
	}

	a := NewAccount(user.FirstName, user.LastName, user.Email,user.Password)
	err = s.store.CreateAccount(a)
	if err!=nil{
		return err 
	}

	tokenStr,err := createJWT(a)
	fmt.Println(tokenStr,"token")
	if err!=nil{
		return err
	}

	
	r.Body.Close()
	return WriteJSON(w, http.StatusOK, a)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	log.Println("Delete called")
	return nil
}

func (s *APIServer) handleAccountByID(w http.ResponseWriter, r *http.Request) error {

	id := mux.Vars(r)["id"]
	if r.Method=="GET"{
		account,err := s.store.GetAccountByID(id)
		if err!=nil{
			return err
		}
		WriteJSON(w,http.StatusOK,account)
		return nil
	}
	if r.Method=="DELETE"{
		return s.store.DeleteAccountByID(id)
	}
	return fmt.Errorf("method not allowed::%s", r.Method)
}


func (s *APIServer) handleGetAccounts(w http.ResponseWriter,r *http.Request)error{
	accounts,err := s.store.GetAccounts()
	if err!=nil{
		return err
	}
	return WriteJSON(w,http.StatusOK,accounts)
	
}


func (s *APIServer) handleTransfer(w http.ResponseWriter,r *http.Request)error{
	transferReq := TransferReq{}
	if err:=json.NewDecoder(r.Body).Decode(&transferReq);err!=nil{
		return err
	}
	defer r.Body.Close()
	return WriteJSON(w,http.StatusOK,transferReq)
}


func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request)error{
	if r.Method != "POST"{
		return fmt.Errorf("method not allowed::%s", r.Method)
	}
	var req LoginReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err!=nil{
		return err 
	}
	acc,err := s.store.GetAccountByEmail(req.Email)
	if err!=nil{
		return err 
	}
	err = bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(req.Password))
	if err!=nil{
		return fmt.Errorf("email or password is wrong")
	}


	// token,err := createJWT(acc)
	
	return nil


}


//middlewares


func createJWT(account *Account)(string,error){
	claims := jwt.MapClaims{
		"expiresAt":15000,
		"userId":account.ID,
		"accountNumber":account.Number,

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}


func withJWTAuth(handleFunc http.HandlerFunc)http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){

		//catch the token in the header and then validate it 
		
		tokenStr := r.Header.Get("x-jwt-token")
		log.Println("TOKEN",tokenStr)
		token,err := validateJWT(tokenStr)
		if err!=nil{
			log.Println(err)
			WriteJSON(w,http.StatusForbidden,ApiError{Error: err.Error()})
			return 
		}

		id := mux.Vars(r)["id"]
		claims,ok := token.Claims.(jwt.MapClaims)
		if !ok{
			accessDenied(w)
			return
		}
		if claims["userId"].(string) != id{
			accessDenied(w)
			return
		}
		handleFunc(w,r)
	}
}


func validateJWT(tokenString string)(*jwt.Token,error){
	secret := os.Getenv("JWT_SECRET")
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
	
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
}


func accessDenied(w http.ResponseWriter){
	WriteJSON(w,http.StatusForbidden,ApiError{Error: "access denied"})	
}