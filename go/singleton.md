### 1. Check-Lock-Check 
See [Marcio's Blog](http://marcio.io/2015/07/singleton-pattern-in-go/)
```go
    type singleton struct {
    }
  
    var instance *singleton
    var lock sync.Mutex
  
    func GetInstance() *singleton {
        // first checking
        if instance == nil {   
            // lock
            lock.Lock()
            // unlock eventually
            defer lock.Unlock()

            // second checking
            if instance == nil {
                instance = &singleton{}
            }
        }
        
        return instance
    }
```

### 2. Use atomic checking
Improve checking with atomic operation
```go
    ...
    var initialized uint32
    var lock sync.Mutex
  
    func GetInstance() *singleton {
        // first checking
        if atomic.LoadUInt32(&initialized) == 1 {
            return instance
        }
    
        // lock
        lock.Lock()
        // unlock eventually
        defer lock.Unlock()

        // second checking
        if initialized == 0 {
            instance = &singleton{}
            atomic.storeUint32(&initialized, 1)
        }
    }
    
    return instance
}
```

### 3. Use standard library
```Once``` is an object that will perform exactly one action.
```go
    var once sync.Once
  
    func GetInstance() *singleton {
        once.Do(func() {
            instance = &singleton{}
        }
    }
    
    return instance
}
