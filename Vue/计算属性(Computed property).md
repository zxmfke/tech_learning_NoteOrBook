计算属性
在模板中绑定表达式是非常便利的，但是它们实际上只用于简单的操作。在模板中放入太多的逻辑会让模板过重且难以维护 所以需要计算属性

#基础例子

```html
<div id="example">
<p>Original message: "{{ message }}"</p>
<p>Computed reversed message: "{{ reversedMessage }}"</p>
</div>

var vm = new Vue({
el: '#example',
data: {
message: 'Hello'
},
computed: {
// a computed getter
reversedMessage: function () {
// `this` points to the vm instance
return this.message.split('').reverse().join('')
}
}
})
```

这里我们声明了一个计算属性 reversedMessage 。我们提供的函数将用作属性 vm.reversedMessage 的 getter function。

```html
console.log(vm.reversedMessage) // -> 'olleH'
vm.message = 'Goodbye'
console.log(vm.reversedMessage) // -> 'eybdooG'
vs.reversedMessage 依赖于vm.message。
```

#计算缓存 vs Methods
通过调用表达式中的method来达到同样的效果

```html
<p>Reversed message: "{{ reverseMessage() }}"</p>

// in component
methods: {
reverseMessage: function () {
return this.message.split('').reverse().join('')
}
}
```

不经过计算属性，我们可以在 method 中定义一个相同的函数来替代它。对于最终的结果，两种方式确实是相同的。然而，不同的是计算属性是基于它的依赖缓存。计算属性只有在它的相关依赖发生改变时才会重新取值。这就意味着只要 message 没有发生改变，多次访问 reversedMessage 计算属性会立即返回之前的计算结果，而不必再次执行函数。关键在于不需要重复函数计算。

#计算属性 vs Watched Property
Vue.js 提供了一个方法 $watch ，它用于观察 Vue 实例上的数据变动。当一些数据需要根据其它数据变化时， $watch很诱人 —— 特别是如果你来自 AngularJS 。不过，通常更好的办法是使用计算属性而不是一个命令式的 $watch 回调。

```html
<div id="demo">{{ fullName }}</div>

var vm = new Vue({
el: '#demo',
data: {
firstName: 'Foo',
lastName: 'Bar',
fullName: 'Foo Bar'
},
watch: {
firstName: function (val) {
this.fullName = val + ' ' + this.lastName
},
lastName: function (val) {
this.fullName = this.firstName + ' ' + val
}
}
})
```

上面代码是命令式的和重复的。跟计算属性对比：

```html
var vm = new Vue({
el: '#demo',
data: {
firstName: 'Foo',
lastName: 'Bar'
},
computed: {
fullName: function () {
return this.firstName + ' ' + this.lastName
}
}
})
```

计算属性也可以放变数，不一定要function。

#计算setter

```html
computed: {
fullName: {
// getter
get: function () {
return this.firstName + ' ' + this.lastName
},
// setter
set: function (newValue) {
var names = newValue.split(' ')
this.firstName = names[0]
this.lastName = names[names.length - 1]
}
}
}
```

如果有其他动作的话，setter顺便赋值。



Watcher（还不太懂，之后仔细看下，如果用到的话）
虽然计算属性在大多数情况下更合适，但有时也需要一个自定义的 watcher 。这是为什么 Vue 提供一个更通用的方法通过 watch 选项，来响应数据的变化。当你想要在数据变化响应时，执行异步操作或开销较大的操作，这是很有用的。
