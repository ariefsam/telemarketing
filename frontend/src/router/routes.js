
const routes = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/Index.vue') },
      { path: '/dashboard', component: () => import('pages/Dashboard.vue') },

      { path: '/customer', component: () => import('pages/Customer.vue') },
      { path: '/telemarketer', component: () => import('pages/Telemarketer.vue') },
      { path: '/admin-customer', component: () => import('pages/AdminCustomer.vue') },
      { path: '/import-customer', component: () => import('pages/ImportCustomer.vue') },
    ],
    meta: {
      requiresAuth: false
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
