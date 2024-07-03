# Concurreny

### This means to do multiple tasks almost at the same time . (but not at the same time , it seems like same time)

### waitgroup : wg.Add(n) --> Here n is the number of go routines
### wg.Wait() --> Does not allow to complete the process
### wg.Done() --> After each task is complete this should be written to mark it as complete

### Mutex are used to lock any critical resource . 
### Channels are a way that allows the go routines to talk with each other.