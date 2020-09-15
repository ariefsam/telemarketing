
const routes = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/Index.vue') },
      { path: '/dashboard', name: 'dashboard', component: () => import('pages/Dashboard.vue') },
      { path: '/customer', name: 'customer', component: () => import('pages/customer/Index.vue') },
      { path: '/customer/:phoneNumber', name: 'customer-detail', component: () => import('pages/customer/Detail.vue') },
      { path: '/call-log', name: 'call-log', component: () => import('pages/call-log/Index.vue') },
      { path: '/report', name: 'report', component: () => import('pages/report/Index.vue') },
      { path: '/telemarketer', name: 'telemarketer', component: () => import('pages/telemarketer/Index.vue') },
      { path: '/telemarketer/create', name: 'telemarketer-create', component: () => import('pages/telemarketer/Create.vue') },
      { path: '/admin-customer', name: 'admin-customer', component: () => import('pages/admin-customer/Index.vue') },
      { path: '/import-customer/:source', name: 'import-customer', component: () => import('pages/admin-customer/ImportCustomer.vue') },
    ],
    meta: {
      requiresAuth: true
    }
  },

  {
    path: "/login",
    name: "login",
    component: () => import("pages/account/Login.vue")
  },

  {
    path: '*',
    component: () => import('pages/Error404.vue')
  }
]

export default routes
