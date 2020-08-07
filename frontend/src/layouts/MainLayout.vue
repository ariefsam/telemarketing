<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="leftDrawerOpen = !leftDrawerOpen"
        />

        <q-toolbar-title>Quasar App</q-toolbar-title>

        <div>
          {{userEmail}}
          <q-btn :to="'/login'" v-show="!isLoggedIn">Login</q-btn>
        </div>
      </q-toolbar>
    </q-header>

    <q-drawer v-model="leftDrawerOpen" show-if-above bordered content-class="bg-grey-1">
      <q-list>
        <q-item-label header class="text-grey-8">Essential Links</q-item-label>
        <EssentialLink v-for="link in essentialLinks" :key="link.title" v-bind="link" />
      </q-list>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script>
import EssentialLink from "components/EssentialLink.vue";

const linksData = [
  {
    title: "Customer",
    caption: "customer list",
    icon: "school",
    link: "/customer",
  },
  {
    title: "Telemarketer",
    caption: "telemarketer admin",
    icon: "school",
    link: "/telemarketer",
  },
  {
    title: "Admin Customer",
    caption: "Admin customer list",
    icon: "chat",
    link: "/admin-customer",
  },
  // {
  //   title: 'Assign New Customer',
  //   caption: 'github.com/quasarframework',
  //   icon: 'code',
  //   link: 'https://github.com/quasarframework'
  // },
  // {
  //   title: 'Discord Chat Channel',
  //   caption: 'chat.quasar.dev',
  //   icon: 'chat',
  //   link: 'https://chat.quasar.dev'
  // },
  // {
  //   title: 'Forum',
  //   caption: 'forum.quasar.dev',
  //   icon: 'record_voice_over',
  //   link: 'https://forum.quasar.dev'
  // },
  // {
  //   title: 'Twitter',
  //   caption: '@quasarframework',
  //   icon: 'rss_feed',
  //   link: 'https://twitter.quasar.dev'
  // },
  // {
  //   title: 'Facebook',
  //   caption: '@QuasarFramework',
  //   icon: 'public',
  //   link: 'https://facebook.quasar.dev'
  // },
  // {
  //   title: 'Quasar Awesome',
  //   caption: 'Community Quasar projects',
  //   icon: 'favorite',
  //   link: 'https://awesome.quasar.dev'
  // }
];

export default {
  name: "MainLayout",
  components: { EssentialLink },
  data() {
    return {
      isLoggedIn: false,
      userEmail: "",
      leftDrawerOpen: false,
      essentialLinks: linksData,
    };
  },
  mounted() {
    var vm = this;

    vm.isLoggedIn = vm.$authService.isTokenValid();
    console.log("Loggedin", vm.isLoggedIn);
    if (vm.isLoggedIn == true) {
      // vm.userName = vm.$authService.getUserName();
      var user = vm.$authService.getUser();
      vm.userEmail = user.Email;
      console.log(user, "ada user");
    }
  },
};
</script>
