import { LocalStorage } from "quasar";
import JwtDecode from 'jwt-decode';
import Vue from 'vue'

var deviceID
var authService = {
  getDeviceID () {
    deviceID = LocalStorage.getItem("device_id");
    if (!deviceID) {
      deviceID = uuidv4()
      LocalStorage.set("device_id", deviceID);
    }
    return deviceID;
  },
  uuidv4 () {
    return ([1e7] + -1e3 + -4e3 + -8e3 + -1e11).replace(/[018]/g, c =>
      (
        c ^
        (crypto.getRandomValues(new Uint8Array(1))[0] & (15 >> (c / 4)))
      ).toString(16)
    );
  },
  setToken (token) {
    LocalStorage.set(this.tokenName(), token);
  },
  getToken () {
    return LocalStorage.getItem(this.tokenName());
  },
  isTokenValid () {
    var currentToken = this.getToken();
    if (!currentToken) {
      return false;
    }
    var auth = JwtDecode(currentToken);
    if (!auth.exp) {
      return false;
    }
    if (auth.exp <= (Date.now() / 1000)) {
      return false
    } else {
      return true
    }
  },
  getUserName () {
    var currentToken = this.getToken();
    if (!currentToken) {
      return "";
    }
    var auth = JwtDecode(currentToken);
    return auth.Telemarketer.Name
  },
  getUser () {
    var currentToken = this.getToken();
    if (!currentToken) {
      return {};
    }
    var auth = JwtDecode(currentToken);
    return auth.Telemarketer
  },
  getPhoneNumber () {
    var user = this.getUser();
    return user.PhoneNumber;
  },
  getRole () {
    var user = this.getUser();
    return user.Role;
  },
  getClinicName () {
    var user = this.getUser();
    return user.ClinicName;
  },
  getProfilePicture () {
    var user = this.getUser();
    return user.Picture;
  },
  logout () {
    LocalStorage.remove(this.tokenName());
  },
  tokenName () {
    return "royalCanin";
  },
  setEmail (email) {
    LocalStorage.set("email", email);
  },
  getEmail () {
    return LocalStorage.getItem("email");
  }
}
// alert(authService);

Vue.prototype.$authService = authService

export default ({ Vue, router }) => {
//   router.beforeEach((to, from, next) => {
//     var isTokenValid = authService.isTokenValid();
//     const requiresAuth = to.matched.some(record => record.meta.requiresAuth);

//     var userDataComplete = false;
//     var phoneNumber = authService.getPhoneNumber();
//     var role = authService.getRole();
//     var clinicName = authService.getClinicName();

//     if (phoneNumber == "" || role == "" || clinicName == "") {
//       userDataComplete = false
//     } else {
//       userDataComplete = true
//     }

//     if (requiresAuth && !isTokenValid) {
//       next('/login');
//     } else {
//       // if user go to /login but he's/she's still autenticated then go to /
//       if (to.path == '/login' && isTokenValid) {
//         next('/')
//       } else {
//         // if user go to / but the profile not completed yet then go to /edit-profile
//         if (to.path == '/' && !userDataComplete) {
//           router.replace({ name: 'edit-profile' })
//         } else {
//           next()
//         }
//       }
//     }
//   });
}
