
const routes = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', name: 'root', component: () => import('pages/Index.vue') },
      { path: '/dashboard', name: 'dashboard', component: () => import('pages/Dashboard.vue') },
      { path: '/admin-dashboard', name: 'admin-dashboard', component: () => import('pages/dashboard/Admin.vue') },
      { path: '/telemarketer-dashboard', name: 'telemarketer-dashboard', component: () => import('pages/dashboard/Telemarketer.vue') },
      { path: '/customer', name: 'customer', component: () => import('pages/customer/Index.vue') },
      { path: '/customer/add', name: 'customer-add', component: () => import('pages/customer/Add.vue') },
      { path: '/customer/:phoneNumber', name: 'customer-detail', component: () => import('pages/customer/Detail.vue') },
      { path: '/closed-customer', name: 'closed-customer', component: () => import('pages/customer/Closed.vue') },
      { path: '/call-log', name: 'call-log', component: () => import('pages/call-log/Index.vue') },
      { path: '/report', name: 'report', component: () => import('pages/report/Index.vue') },
      { path: '/telemarketer', name: 'telemarketer', component: () => import('pages/telemarketer/Index.vue') },
      { path: '/telemarketer/create', name: 'telemarketer-create', component: () => import('pages/telemarketer/Create.vue') },
      { path: '/telemarketer/:id/edit', name: 'telemarketer-edit', component: () => import('pages/telemarketer/Edit.vue') },
      { path: '/admin-customer', name: 'admin-customer', component: () => import('pages/admin-customer/Index.vue') },
      { path: '/admin-customer/:phoneNumber', name: 'admin-customer-detail', component: () => import('pages/admin-customer/Detail.vue') },
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
