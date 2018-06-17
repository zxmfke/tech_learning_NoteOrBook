Class 与 Style 绑定
数据绑定一个常见需求是操作元素的 class 列表和它的内联样式。因为它们都是属性 ，我们可以用v-bind 处理它们：只需要计算出表达式最终的字符串。不过，字符串拼接麻烦又易错。因此，在 v-bind 用于 class 和 style 时， Vue.js 专门增强了它。表达式的结果类型除了字符串之外，还可以是对象或数组。

绑定HTML Class
#对象语法
我们可以传给 v-bind:class 一个对象，以动态地切换 class 。

```html
<div v-bind:class="{ active: isActive }"></div>
```

你也可以直接绑定数据里的一个对象

```html
<div v-bind:class="classObject"></div>
data: {
classObject: {
active: true,
'text-danger': false
}
}
```

我们也可以在这里绑定返回对象的计算属性

```html
<div v-bind:class="classObject"></div>
data: {
isActive: true,
error: null
},
computed: {
classObject: function () {
return {
active: this.isActive && !this.error,
'text-danger': this.error && this.error.type === 'fatal',
}
}
}
```

#数组语法
我们可以把一个数组传给 v-bind:class ，以应用一个 class 列表

```html
<div v-bind:class="[activeClass, errorClass]">
data: {
activeClass: 'active',
errorClass: 'text-danger'
}
```

如果你也想根据条件切换列表中的 class ，可以用三元表达式：

```html
<div v-bind:class="[isActive ? activeClass : '', errorClass]">
```

#用在组件上
当你在一个定制的组件上用到 class 属性的时候，这些类将被添加到根元素上面，这个元素上已经存在的类不会被覆盖。组件也适用于绑定HTML Class

```html
Vue.component('my-component', {
template: '<p class="foo bar">Hi</p>'
})
<my-component class="baz boo"></my-component>
```

HTML 最终将被渲染成为:

```html
<p class="foo bar baz boo">Hi</p>
```

绑定内联样式
#对象语法
v-bind:style 的对象语法十分直观——看着非常像 CSS ，其实它是一个 JavaScript 对象。 CSS 属性名可以用驼峰式（camelCase）或短横分隔命名（kebab-case）：

```html
<div v-bind:style="{ color: activeColor, fontSize: fontSize + 'px' }"></div>
data: {
activeColor: 'red',
fontSize: 30
}
```

直接绑定到一个样式对象通常更好，让模板更清晰：

```html
<div v-bind:style="styleObject"></div>
data: {
styleObject: {
color: 'red',
fontSize: '13px'
}
}
```

#数组语法
和绑定HTML Class是一样的，可以用数组bind v-bind:style

```html
#Multiple Values
Starting in 2.3 you can provide an array of multiple (prefixed) values to a style property
<div v-bind:style="{ display: ['-webkit-box', '-ms-flexbox', 'flex'] }">
```






