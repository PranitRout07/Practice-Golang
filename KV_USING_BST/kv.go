package main

import "fmt"

type KeyVal struct {
	key int
	value string 
}

type KeyValueDB struct {
	KeyVal 
	left  *KeyValueDB
	right *KeyValueDB
}



func CreateRoot(key int,value string) *KeyValueDB {

	return &KeyValueDB{
		KeyVal: KeyVal{key,value},
	}
}

func (b *KeyValueDB) Insertion(key int,value string) *KeyValueDB {
	
		
		newNode := &KeyValueDB{KeyVal: KeyVal{key,value}}

	if b==nil{
		return newNode
	}

	if b.key > newNode.key {
		b.left = b.left.Insertion(key,value)
		return b
	} 
	b.right = b.right.Insertion(key,value)
	return b
}

func (b *KeyValueDB) SearchElement(key int) bool{
	if b==nil{
		
		return false;
	}

	if b.key==key{
		
		return true;
	}

	if key>=b.key{
		return true && b.right.SearchElement(key)
	}else{
		return true && b.left.SearchElement(key)

	}

}

func ( b *KeyValueDB)PrintAllLeafNodes(){
	if b==nil{
		return
	}

	if b.left==nil && b.right==nil{
		fmt.Print(" ",b.value," ")
	}
	b.left.PrintAllLeafNodes()
	b.right.PrintAllLeafNodes()
}


func (b *KeyValueDB) PreorderTraverse(){
	if b==nil{
		return
	}
	fmt.Print(" ",fmt.Sprintf("%d:%s",b.key,b.value)," ")
	b.left.PreorderTraverse()
	b.right.PreorderTraverse()
}	

func (b *KeyValueDB) InorderTraverse(){
	if b==nil{
		return
	}
	b.left.PreorderTraverse()
	fmt.Print(" ",fmt.Sprintf("%d:%s",b.key,b.value)," ")
	b.right.PreorderTraverse()
}

func (b *KeyValueDB)MinNode()*KeyValueDB{
	if b==nil || b.left==nil{
		return b 
	}
	
	return b.left.MinNode()
}
func (b *KeyValueDB)MaxNode()*KeyValueDB{
	if b==nil || b.right==nil{
		return b 
	}
	
	return b.right.MaxNode()
 
}

func (b *KeyValueDB)DeleteElement(key int) *KeyValueDB{
	if b==nil{
		fmt.Println("Key not present")
		return nil 
	}
	if b.key==key{
		
	
		if b.left==nil && b.right==nil{
			b = nil 
			return b 
		}else if b.left==nil && b.right!=nil{
			b = b.right 
			return b 
		}else if b.left!=nil && b.right==nil{
			b = b.left
			return b
		}else if (b.left!=nil && b.right!=nil){
			//max find from left
			//min find from right 
			data := b.right.MinNode()
			b.key = data.key
			b.value = data.value
			b.right = b.right.DeleteElement(data.key)

			return b
		}
	}

	if b.key<=key{
		b.right = b.right.DeleteElement(key)
	}else{
		b.left = b.left.DeleteElement(key)
	}
	return b

}

// func (b *KeyValueDB) IsBalanced() bool{
// 	if 
// }

func (b *KeyValueDB) GetHeight(key int)int{
	count:=0

	if b==nil{
		return 0
	}

	if !b.SearchElement(key){
		fmt.Println("No such element")
		return -1
	}
	if b.key == key{
		return count
	}
	
	if key < b.key{
		count  = count +1

		return count + b.left.GetHeight(key) 
	}
	count  = count +1

	

	return count + b.right.GetHeight(key) 
	
}


func main() {
	root := CreateRoot(20,"demo")
	fmt.Println(root)
	root.Insertion(15,"Hii")
	root.Insertion(17,"Hello")
	root.Insertion(31,"Greater")
	root.Insertion(100,"Greatest")
	root.InorderTraverse()
	// root = root.Insertion(10)
	// root  = root.Insertion(31)
	// root = root.Insertion(19)
	// root = root.Insertion(21)
	// root.PreorderTraverse()
	// fmt.Println(root.SearchElement(19))
	// root = root.DeleteElement(20)
	// root = root.DeleteElement(31)
	// root = root.DeleteElement(21)
	// root = root.DeleteElement(10)
	// root = root.DeleteElement(19)
	// fmt.Println()
	// root.PrintAllLeafNodes()
	// fmt.Println()
	// root.PreorderTraverse()
	// fmt.Println(root.GetHeight(19))
	// for {
	// 	var input int
	// 	fmt.Println()
	// 	fmt.Println("Enter value:-")
	// 	fmt.Scanln(&input)
	// 	root = root.Insertion(input)



	// 	root.PreorderTraverse()


	// 	var choice string
	// 	fmt.Println()
	// 	fmt.Println("Do You Want To Search An Element From The Binary Tress?")
	// 	fmt.Scanln(&choice)
	// 	if choice=="yes"{
	// 		var val int 
	// 		fmt.Println("Enter element:-")
	// 		fmt.Scanln(&val)
	// 		root.SearchElement(val)
	// 	}

	// }
}
