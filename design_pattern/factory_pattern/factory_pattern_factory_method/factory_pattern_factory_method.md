### 工厂方法模式(Factory Method Pattern)

---

<font size="5">**工厂方法模式**</font>定义了一个创建对象的接口，但由子类决定要实例化的类是哪一个。工厂方法让类把实例化推迟到子类。

---

**看个例子**

在下面的例子中，我们有一个pizza店，有一个订单系统。每一个pizza都要经过准备，烘焙，切片和装盒。我们的订单系统依赖于每个不同风味的pizza，每次都要实现对应的pizza之后才开始制作。然而我们不知道我们要制作的pizza是哪一种风味，因此它不需要也不应该和具体的要制作的pizza对象耦合在一起。

```go
func (p *Pizza) OrderPizza(typ,store string) *Pizza{
    var pizza *Pizza = nil
    // 根据pizza的类型和store类型，我们实例化正确的具体类，然后将其赋值给pizza实例变量。请注意，这里的任何pizza都是实现Pizza接口。
    if store == "NY" {
        switch (typ) {
    	case "cheese" :
            pizza = new(NewYorkCheesePizza)
        	break
    	case "greek" :
        	pizza = new(NewYorkGreekPizza)
        	break
    	case "pepperoni" :
        	pizza = new(NewYorkPepperoniPizza)
        	break    
    	...
        }
    } else if store == "XM" {
        switch (typ) {
    	case "cheese" :
        	pizza = new(XiaMenCheesePizza)
        	break
    	case "greek" :
        	pizza = new(XiaMenGreekPizza)
        	break
    	case "pepperoni" :
            pizza = new(XiaMenPepperoniPizza)
        	break    
    	...
        }
    }
    
    pizza.prepare()
    pizza.bake()
    pizza.cut()
    pizza.box()
    
    return pizza
}
```



---

**###设计原则**

> **依赖倒置原则：要依赖抽象，不要依赖具体类。不能让高层组件依赖底层组件，而且，不管高层或低层组件，“两者都应该依赖于抽象”**

将上面的例子通过依赖倒置原则转换一下，可以看到pizzaStore不再依赖于各个pizza对象，而是依赖于一个pizza抽象类，而下面每种风味的pizza也继承与pizza抽象类。这个就是所谓依赖倒置。



---

**意图**

定义一个创建对象的接口，但是让它的子类决定实例哪种类型，延迟到子类来选择实现。工厂方法模式的抓哟功能是让父类在不知道具体实现的情况下，完成自身的功能调用。

---

**组成部分**

- AbstractFactory (抽象创建对象的工厂接口，java用Abstract)

  一个接口或一个类，负责定义一个创建对象的抽象方法，该方法返回具体产品类的实例

- ConcreteFactory(具体实现封装创建对象的工厂)

  具体工厂，决定如何实例化具体产品，需要有多少种具体产品，就有多少个工厂。

- Product (工厂生产的对象的超类)

  一个实现要被生产的对象的接口方法的超类，定义一些产品必须实现的方法

- ConcreteProduct(具体的工厂要生产的对象)

  具体实现实例化的方法

![](.\image\classic_factory_pattern_class_diagram.jpg)

---

**范例代码**

代码：http://192.168.1.189/navi/wiki/tree/master/msa/design_pattern/factory_pattern/factory_pattern_factory_method/example

范例代码为一个商店卖汉堡，商店的接口定义了下单接口，Mcdonald实现下单接口。下单方法里面通过typ string来决定要实例化的对象，也就是哪一种汉堡。汉堡接口定义了Prepare(),Bake(),Package()和GetName()方法，然后超类Hamburger实现了这四个方法，即实现了汉堡接口。DoubleCheeseBurger，BigMac和Unworthy分别继承了Hamburger超类，即也实现了汉堡接口。

在商店的OrderHamburger中，实例化的对象是实现了汉堡接口的对象，所以都可以调用接口的四个方法。所以在下方代码中，不需要知道具体的子类是谁，只需要知道实例的是一个汉堡，有实现Prepare(),Bake(),Package()和GetName()接口就可以了。

```go
func (m *Moss) OrderHamburger(typ string) HamburgerBasic {
	hamburger := m.CreateHamburger(typ)
	hamburger.Prepare()
	hamburger.Bake()
	hamburger.Package()
	fmt.Printf("This is your %s\n", hamburger.GetName())
	fmt.Println("--------------------------------------")
	return hamburger
}

func (m *Moss) CreateHamburger(typ string) HamburgerBasic {
	switch (typ) {
	case "Double Cheeseburger":
		return NewMossDoubleCheeseBurger()
	case "Big Mac":
		return NewMossBigMac()
	case "Unworthy":
		return NewMossUnworthy()
	default:
		return nil
	}
}
```



![](.\image\classic_factory_pattern_class_diagram_v1.jpg)

---

**总结**

- 在编程中，产品类的实例化有时候是比较复杂和多变的，通过工厂模式，将产品的实例化封装起来，使得调用者根本无需关心产品的实例化过程，只需依赖工厂即可得到自己想要的产品。
- 作为一种创建类模式，在任何需要生成复杂对象的地方，都可以使用工厂方法模式。有一点需要注意的地方就是复杂对象适合使用工厂方法模式，而简单对象，特别是只需要通过new就可以完成创建的对象，无需使用工厂方法模式。如果使用工厂方法模式，就需要引入一个构造者类，会增加系统的复杂度。
- 工厂方法模式是一种典型的解耦模式，如果调用者自己组装产品需要增加依赖关系时，可以考虑使用工厂方法模式，将会大大降低对象之间的耦合度。
- 由于工厂方法模式是依靠抽象架构的，它把实例化产品的任务交由实现类完成，扩展性比较好。也就是说，当需要系统有比较好的扩展性时，可以考虑工厂模式，不同的产品用不同的具体构造者组装。

---

**区别**

从本质上讲，简单工厂和工厂方法都是在“选择实现”。但简单工厂是直接在工厂类里面进行“选择实现”；而工厂方法会把这个工作延迟到子类来实现，工厂类里面使用工厂方法的地方是依赖于抽象而不是具体的实现，从而使得系统更加灵活，具有更好的可维护性和可扩展性。

---

资料来源

<https://www.cnblogs.com/dazuihou/p/3639367.html>