### 外观模式(Facade Pattern)

---

<font size="5">**外观模式**</font>提供了一个统一的接口，用来访问子系统中的一群接口。外观定义了一个高层接口，让子系统更容易使用。

---

**看个例子**

一台电脑的开机，需要经过BIOS load，主引导记录MBR，硬盘启动......每一步都是一个复杂的子系统。然而我们的使用者在开机的时候，肯定不会是自己一步一步这样开机，只会按一个开机键。开机键就是封装开机的所有必须经过的步骤。

---

**设计原则**

>**最少知识原则：只和你的密友谈话**

就任何对象而言，在改对象的方法内，我们只应该调用属于以下范围的方法：

- 该对象本身
- 被当做方法的参数而传递进来的对象
- 此方法所创建或实例化的任何对象
- 对象的任何组件

```go
// 没有用原则，这里从气象站取得了温度计对象，然后再从温度计对象取得温度
func getTemp() float64{
    thermonmeter = station.getThermonmeter()
    return thermometer.getTemperature()
}

// 采用原则，这里在气象站 中加进一个方法，用来想温度计请求温度。这可以减少我们所依赖的类的书目
func getTemp() float64{
    return station.getTemperature()
}
```

```go
type Car struct{
    Engine Engine
}

func (c *Car) Start(key Key){ // 被当做参数传进来的对象，其方法可以被调用
    doors := new(Doors) // 在这里创建了一新的对象，它的方法可以被调用
    
    authorized = key.turns()
    
    if authorized {
        c.engine.start() //可以调用对象组件的方法
        c.UpdateDashboardDisplay() //可以调用同一个对象内的方法
        doors.lock() //可以调用你所创建或实例化的对象的方法
    }
}

func (c *Car) UpdateDashboardDisplay() {
    
}
```

---

**意图**

封装监护，简化调用

---

**范例代码**

- 类的方法实现外观模式

```go
type CPU struct {}

func (this *CPU) Freeze() {
    fmt.Println("CPU frozen")
}

func (this *CPU) Jump(position int64) {
	fmt.Println("CPU jumps to", position)
}

func (this *CPU) Execute() {
    fmt.Println("CPU executes")
}

type Memory struct {}

func (this *Memory) Load(position int64, data []byte) {
	fmt.Println("Memory loads")
}

type HardDrive struct {}

func (this *HardDrive) Read(lba int64, size int64) []byte {
	fmt.Println("HardDrive reads")
	return nil
}

type Computer struct{
	cpu *CPU
	memory *Memory
	hardDrive *HardDrive
}

func NewComputer() *Computer {
	cpu := new(CPU)
	memory := new(Memory)
	hardDrive := new(HardDrive)
	return &Computer{cpu,memory,hardDrive}
}

func (this *Computer) Start() {
	this.cpu.Freeze()
	this.memory.Load(0, this.hardDrive.Read(0,1023))
	this.cpu.Jump(10)
	this.cpu.Execute()
}

func main() {
    computer := NewComputer()
	computer.Start()
	//也可以调用子系统的方法
	subpart := new(CPU)
	subpart.Execute()
}
```

- 包的方法实现外观模式

```go
package computer

import "system"

var cpu *system.CPU
var memory *system.Memory
var hardDrive *system.HardDrive

func init() {
	cpu = new(system.CPU)
	memory = new(system.Memory)
	hardDrive = new(system.HardDrive)
}

func Start() {
	cpu.Freeze()
	memory.Load(0, hardDrive.Read(0,1023))
	cpu.Jump(10)
	cpu.Execute()
}
```

```go
import "computer"
...
computer.Start()
```

---

**总结**

- 外观模式的目的不是给子系统添加新的功能接口，而是为了让外部减少与子系统内多个模块的交互，松散耦合，从而让外部能否更简单的使用子系统。
- 外观模式让子系统更加易用，客户端不再需要了解子系统内部的实现，也不需要跟众多子系统内部的模块进行交互，只需要跟外观交互就可以了，相当于外观类为外部客户端使用子系统体统了一站式服务。
- 把需要暴露给外部的功能集中到外观中个，这样既方便客户端使用，也很好地隐藏内部的细节。

