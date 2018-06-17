构造函数
每个 Vue ViewModel 都是通过 Vue 构造函数创建出一个 Vue 根实例来引导辅助的

```html
var vm = new Vue({
// 选项
})
```

在实例化 Vue 实例时，你需要传入一个选项对象，它可以包含数据(data)、模板(template)、挂载元素(element to mount on)、方法(methods)、生命周期函数(lifecycle callbacks)和其他选项。

可以通过预先定义选项扩展 Vue 构造函数，从而创建可复用的组件构造函数：

```html
var MyComponent = Vue.extend({
// 扩展选项
})
// `MyComponent` 的所有实例，都将由预先定义的扩展选项来创建
var myComponentInstance = new MyComponent()
```

属性和方法

每个 Vue 实例都会代理其 data 对象的所有属性

```html
var data = { a: 1 }
var vm = new Vue({
data: data
})
vm.a === data.a // -> true
// 设置属性也会影响到原始数据
vm.a = 2
data.a // -> 2
// ... 反之亦然
data.a = 3
vm.a // -> 3
```

应当注意，只有这些以上这种代理属性是响应式的。如果在实例创建之后，再对实例添加新的属性，将不会触发任何视图更新。在建立instance之前赋值是互相影响的，如果是在new之后对object修改是不糊影响的。

钩子函数

每个 Vue 实例在被创建之前，都要经过一系列的初始化过程 - 例如，Vue 实例需要配置数据观察(data observation)、编译模版(compile the template)、在 DOM 挂载实例(mount the instance to the DOM)，以及在数据变化时更新 DOM(update the DOM when data change)。在这个过程中，Vue 实例还会调用执行一些生命周期钩子函数，这就为我们提供了执行自定义逻辑的时机。

```html
var vm = new Vue({
data: {
a: 1
},
created: function () {
// `this` 指向 vm 实例
console.log('a is: ' + this.a)
}
})
```

钩子函数比较像在建立实例时的一个debug和log的动作，管理这个实例。

实例生命周期

![Image](https://github.com/zxmfke/tech_learning_NoteOrBook/blob/master/Vue/assets/Image1.png)
