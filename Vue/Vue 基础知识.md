Vue.js是一套构建用户界面(user interface)的渐进式框架。与其他重量级框架不同的是，Vue 从根本上采用最小成本、渐进增量(incrementally adoptable)的设计。Vue 的核心库只专注于视图层，并且很容易与其他第三方库或现有项目集成。
======================
声明式渲染
Vue.js 的核心是，可以采用简洁的模板语法来声明式的将数据渲染为 DOM

```html
<div id="app">
  {{ message }}
</div>
```

======================
另一个方法绑定DOM

```html
<div id="app-2">
  <span v-bind:title="message">
    鼠标悬停此处几秒，
    可以看到此处动态绑定的 title！
  </span>
</div>
```

这里我们遇到一些新内容。你看到的 v-bind 属性被称为指令。指令带有前缀 v-，表示是由 Vue 提供的专用属性。可能你已经猜到了，它们会在渲染的 DOM 上产生专门的响应式行为。简而言之，这里该指令的作用就是：“将此元素的 title 属性与 Vue 实例的 message 属性保持关联更新”。
=======================
on监听事件

```html
<div id="app-5">
<p>{{ message }}</p>
<button v-on:click="reverseMessage">翻转 message</button>
</div>
```

reverseMessage为function不用加()，在这个function中只会改app中的message并不会改动DOM。
=======================
v-model 输入与app之间的双向绑定

```html
<div id="app-6">
<p>{{ message }}</p>
<input v-model="message">
</div>
```

========================
由组件组合而成的应用程序
组件系统是 Vue 的另一个重要概念，因为它是一种抽象，可以让我们使用小型、自包含和通常可复用的组件，把这些组合来构建大型应用程序

Vue中注册组件

```html
Vue.component('todo-item', {
template: '<li>这是一个 todo 项</li>'
})
```

组件在组件模版中应用

```html
<ol>
<!-- 创建一个 todo-item 组件的实例 -->
<todo-item></todo-item>
</ol>
```

为了让组件更加灵活，将数据从父作用域传到子组件，添加属性props

```html
Vue.component('todo-item', {
// todo-item 组件现在接受一个 "prop"，
// 类似于一个自定义属性。
// 此 prop 名为 todo。
props: ['todo'],
template: '<li>{{ todo.text }}</li>'
})
```
