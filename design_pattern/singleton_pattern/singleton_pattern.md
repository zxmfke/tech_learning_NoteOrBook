### 单例模式(Singleton Pattern)

---

<font size="5">**单例模式**</font>确保一个类只有一个实例，并提供一个全局访问点。

---

**动机**

对于系统中的某些类来说，只有一个实例很重要，例如，一个系统中可以存在多个打印任务，但是只能有一个正在工作的任务；一个系统只能有一个窗口管理器或文件系统；一个系统只能有一个计时工具或ID(序号)生成器。如在Windows中就只能打开一个任务管理器。如果不使用机制对窗口对象进行唯一化，将弹出多个窗口，如果这些窗口显示的内容完全一致，则是重复对象，浪费内存资源；如果这些窗口显示的内容不一致，则意味着在某一瞬间系统有多个状态，与实际不符，也会给用户带来误解，不知道哪一个才是真实的状态。因此有时确保系统中某个对象的唯一性即一个类只能有一个实例非常重要。

---

**意图**

保证整个系统只会有一个实例，并提供一个全局访问点

---

**组成部分**

- 一个全局变量
- 一个获取实例的function

---

范例代码

- 最简单的单例模式，非线程安全

  在单线程下，下面代码块确实可以保证全局唯一一个config实例，但是在多线程下未必。因为第二个线程执行的时候，可能刚好在第一个线程执行到第11行的时候，第二个线程也来取实例，那config就会被实例化。然而，config在第一个线程执行的时候也会实例化。那在多线程的情况下，就还是会造成多次实例化的情况。

  ```go
  package main
  
  import fmt
  
  type Config struct{
      
  }
  
  var config *Config
  
  func GetConfig() *Config{
      
      if config == nil {
          config = &Config{}
          fmt.Println("this is first init")
      } else {
          fmt.Println("already init")
      }
      
      return config
  }
  
  func main() {
      config = GetConfig()
      config = GetConfig()
  }
  ```

- 加锁的单例模式，避免多线程造成的问题

  可以通过加锁来避免多线程造成多实例的问题

  ```go
  package main
  
  import (
      "sync"
      "fmt"
  )
  
  type Config struct{
      
  }
  
  var config *Config
  var lock *sync.Mutex = &sync.Mutex {}
  
  func GetConfig() *Config{
      lock.Lock()
      defer lock.Unlock()
      if config == nil {
          config = &Config{}
          fmt.Println("this is first init")
      } else {
          fmt.Println("already init")
      }
      
      return config
  }
  
  func main() {
      config = GetConfig()
      config = GetConfig()
  }
  ```

- 双重锁机制

  上方代码块在每次取实例的时候，都会加锁，解锁，影响效能，所以在外层先做一次判断

  ```go
  package main
  
  import (
      "sync"
      "fmt"
  )
  
  type Config struct{
      
  }
  
  var config *Config
  var lock *sync.Mutex = &sync.Mutex {}
  
  func GetConfig() *Config{
      
      if config == nil {
      	lock.Lock()
      	defer lock.Unlock()
      	if config == nil {
          	config = &Config{}
          	fmt.Println("this is first init")
      	} else {
          	fmt.Println("already init")
      	}
      }
      
      return config
  }
  
  func main() {
      config = GetConfig()
      config = GetConfig()
  }
  ```

- Go自带的标准库sync.Once实现单例模式

  ```go
  package main
  
  import (
      "sync"
      "fmt"
  )
  
  type Config struct{
      
  }
  
  var config *Config
  var once sync.Once
  
  func GetConfig() *Config{
      
      once.Do(func(){
          config = &Config{}
        	fmt.Println("this is first init")
      })
      
      return config
  }
  
  func main() {
      config = GetConfig()
      config = GetConfig()
  }
  ```

---

**总结**

单例模式会阻止其他对象实例化其自己的单例对象的副本，从而确保所有对象都访问唯一实例。