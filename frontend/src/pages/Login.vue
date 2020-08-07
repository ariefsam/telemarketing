<template>
  <q-layout>
    
    <q-page-container class="page-container">
      <q-page class="row justify-center items-center">
        <div class="col-md-5 login-container text-justify">
          <div class="title">Welcome,</div>
          
          <div
            class="row items-center justify-center social-button q-mt-md q-mb-sm"
            @click="loginByGoogle"
          >
            <img src="~assets/google-icon-small.png" width="22" />
            <div class="q-pl-sm">Log In with Google</div>
          </div>
        </div>
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script>
export default {
  name: "Login",

  data() {
    return {
      email: "",
    };
  },

  methods: {
    loginByGoogle() {
      var vm = this;
      vm.$auth
        .signInWithPopup(vm.$googleProvider)
        .then(function(result) {
          var token = result.credential.accessToken;
          var user = result.user;
          user.getIdToken(false).then(function(idToken) {
            var dataSubmit = {
              FirebaseToken: idToken
            };
            vm.$axios
              .post("/api/login/firebase", dataSubmit)
              .then(function(response) {
                if (response.data.Token) {
                  vm.$authService.setToken(response.data.Token);
                  vm.$router.push("/");
                }
              })
              .catch(function(error) {

              });
          });
        })
        .catch(function(error) {
          var errorCode = error.code;
          var errorMessage = error.message;
          var email = error.email;
          var credential = error.credential;
        });
    },
  },
};
</script>

<style lang="scss">
body.screen--xs {
  .logo-rc {
    left: -12px;
    width: 100px;
  }
  .page-container {
    .login-container {
      padding: 20px 30px;
      max-width: 340px;
    }
  }
}

body.screen--sm {
  .logo-rc {
    left: -6px;
  }
}

.logo-rc {
  position: absolute;
  top: 12px;
  left: 12px;
}
.page-container {
  margin: auto;
  background-repeat: no-repeat;
  background-size: cover;
  background-position: center;
  .login-container {
    position: absolute;
    max-width: 420px;
    background-color: rgba(235, 235, 235, 0.4);
    padding: 30px 50px;
    border-radius: 12px;
    -webkit-backdrop-filter: blur(12px);
    backdrop-filter: blur(15px);
    box-shadow: 0 0 20px $grey-5;
    .title {
      font-size: 26px;
      font-weight: bold;
      letter-spacing: 1px;
      line-height: 1.2;
    }
    .subtitle {
      margin-bottom: 10px;
      font-size: 16px;
      color: $grey-6;
    }
    .login-btn {
      width: 100%;
    }
    .separator {
      background-color: #aaa;
    }
    .q-field--filled .q-field__control {
      border-radius: 8px !important;
    }
    .q-btn--rectangle {
      border-radius: 8px;
      padding: 2px 0;
    }
    .social-button {
      width: 100%;
      font-weight: 400;
      font-size: 14px;
      letter-spacing: 0.02em;
      border-radius: 8px;
      cursor: pointer;
      background-color: white;
      padding: 9px 0;
    }
  }
}
</style>