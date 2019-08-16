container-vue 部分

请先参照 [vue](vue.md)安装好模板

该文档的基础目录是 shipping/vue

修改vue.config.js

```js
 devServer: {
    port: port,
    open: true,
    overlay: {
      warnings: false,
      errors: true
    },
    proxy: {
      // change xxx-api/login => mock/login
      // detail: https://cli.vuejs.org/config/#devserver-proxy
      [process.env.VUE_APP_BASE_API]: {
        target: `http://127.0.0.1:8080`,
        changeOrigin: true,
        pathRewrite: {
          ['^' + process.env.VUE_APP_BASE_API]: ''
        }
      }
    }
  }
```

修改.env.development 与　.env.production
```env
VUE_APP_BASE_API = '/api'
```

修改src 下面的permissioni.js文件

```js
// 在beforeEach 函数中加入，跳过权限认证
  next()
  return
```

增加集装箱路由
```js
  {
    path: '/container',
    component: Layout,
    children: [{
      path: 'index',
      name: 'Container',
      component: () => import('@/views/dashboard/index'),
      meta: { title: '集装箱管理', icon: 'dashboard' }
    }]
  },
```