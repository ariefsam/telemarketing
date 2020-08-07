
const routes = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/Index.vue') },
      { path: '/login', component: () => import('pages/Login.vue') },
      { path: '/customer', component: () => import('pages/Customer.vue') },
      { path: '/telemarketer', component: () => import('pages/Telemarketer.vue') },
      { path: '/admin-customer', component: () => import('pages/AdminCustomer.vue') },
      { path: '/import-customer', component: () => import('pages/ImportCustomer.vue') },
    ]
  },

  {
    path: '*',
    component: () => import('pages/Error404.vue')
  }
]

export default routes
