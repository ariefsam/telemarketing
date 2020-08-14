<template>
  <q-layout view="hHh LpR fFf">
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
        <q-toolbar-title>Telemarketing App</q-toolbar-title>
        <q-space />
        <!--here profile icon-->
        <div class="profile-menu">
          <q-btn flat no-caps>
            <q-avatar size="34px">
              <img :src="user.profilePic" @error="onImageLoadFailure($event)"/>
            </q-avatar>
            <div class="name q-ml-sm">{{user.name}}</div>
            <q-menu transition-show="jump-down" transition-hide="jump-up">
              <div style="width: 240px">
                <div class="row no-wrap items-center profile-info">
                  <q-avatar size="34px">
                    <img :src="user.profilePic" @error="onImageLoadFailure($event)"/>
                  </q-avatar>
                  <div class="name">{{user.email}}</div>
                </div>
                <q-separator inset />
                <div class="row no-wrap profile-options">
                  <q-item clickable v-close-popup class="option" @click="logout">
                    <q-item-section avatar style="min-width: 0">
                      <q-icon name="exit_to_app" />
                    </q-item-section>
                    <q-item-section>Keluar</q-item-section>
                  </q-item>
                </div>
              </div>
            </q-menu>
          </q-btn>
        </div>
      </q-toolbar>
    </q-header>

    <q-drawer show-if-above bordered :width="250" content-class="bg-grey-0" v-model="leftDrawerOpen">
      <q-scroll-area class="fit">
        <q-list v-for="(menuItem, index) in menuList" :key="index">
          <q-item clickable :to="menuItem.link" v-ripple>
            <q-item-section avatar>
              <q-icon :name="menuItem.icon" />
            </q-item-section>
            <q-item-section>{{ menuItem.label }}</q-item-section>
          </q-item>

          <q-separator v-if="menuItem.separator" />
        </q-list>
      </q-scroll-area>
    </q-drawer>

    <!--<q-page-container>
      <router-view />
    </q-page-container>-->
  </q-layout>
</template>

<script>
export default {
  name: "MainLayout",
  
  data() {
    return {
      leftDrawerOpen: false,
      menuList: [
        {
          icon: "dashboard",
          label: "Dashboard",
          separator: true,
          link: { name: "dashboard" }
        },
        {
          icon: "fas fa-user",
          label: "Customer",
          separator: false,
          link: { name: "customer" }
        },
        {
          icon: "fas fa-headset",
          label: "Telemarketer",
          separator: false,
          link: { name: "telemarketer" }
        },
        {
          icon: "fas fa-user-tie",
          label: "Admin Customer",
          separator: false,
          link: { name: "admin-customer" }
        }
      ],
      user:{},
    };
  },
  mounted() {
    var vm = this;
    this.user = {
      profilePic: vm.$authService.getProfilePicture(),
      name: vm.$authService.getUserName(),
      email: vm.$authService.getEmail()
    }
    // Dummy
    this.user = {
      profilePic: 'test',
      name: 'Agung Kurniawan',
      email: 'agung@fsn.co.id'
    }
  },

  methods: {
    logout() {
      var vm = this;
      vm.$authService.logout();
      this.$router.replace({ name: 'login' })
    },
    onImageLoadFailure (event) {
      event.target.src = require('assets/default.jpg')
    }
  }
};
</script>
