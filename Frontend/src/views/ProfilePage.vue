<!-- src/views/ProfilePage.vue -->
<template>
  <div>
    <div v-if="LoggedOut" class="login-page">
        <!-- Login Form -->
        <transition name="slide">
          <div v-if="ShowLogin" class="edit-menu">
            <h3>Login</h3>
            <input v-model="LoginId" placeholder="Enter id">
            <div class="edit-menu-buttons">
                <button @click="Login">Log In</button>
            </div>
            <a href="#" class="toggle-link" @click="ToggleLogIn">Don't have an Account? Register</a>
          </div>
        </transition>
        <transition name="slide">
          <div v-if="ShowRegister" class="edit-menu">
            <h3>Register</h3>
            <input v-model="NewProfile.name" placeholder="Enter Username">
            <input v-model="NewProfile.image" placeholder="Enter Profile Picture URL">
            <div class="edit-menu-buttons">
                <button @click="AddProfile">Register</button>
            </div>
            <a href="#" class="toggle-link" @click="ToggleLogIn">Don't have an Account? Register</a>
          </div>
        </transition>
    </div>
    <!-- 
      Displays all profiles and an "Add Profile" button. 
      Clicking the button shows a form to add a new profile using transitions.
    -->
    <div class="profile-container">
  
        <!-- Loop through existing profiles and render each one -->
        <ProfileCard
            v-if="profile"
            :key="profile.profile_id"
            :profile="profile"
            @toggle-edit-profile="toggleEditProfile"
            @log-out="logout"
        />

        <!-- Edit Profile Form -->
        <transition name="slide">
          <div v-if="ShowEditProfile" class="edit-menu">
            <h3>Edit Profile</h3>
            <input v-model="EditProfile.name" placeholder="Enter Username">
            <input v-model="EditProfile.image" placeholder="Enter Profile Picture URL">
            <div class="edit-menu-buttons">
                <button @click.stop="cancelEdit">Cancel</button>
                <button @click.stop="editProfile">Edit</button>
            </div>
          </div>
        </transition>
    </div>
  </div>
</template>

<script>
  import axios from 'axios'
  import ProfileCard from '@/components/ProfileCard.vue'

  export default {
    name: 'ProfilePage',
    components: { ProfileCard },
    data: function() {
      return {
        ShowAddProfile: false, // Controls visibility of the "Add Profile" modal
        ShowLogin: true,
        ShowRegister: false,
        ShowEditProfile: false,
        NewProfile: {
          name: "",
          image: "",
        },
        LoginId: null,
        EditProfile: {
          id: null,
          name: "",
          image:"",
        },
      };
    },
    computed: {
      profile() {
        return this.$root.profile;
      },
      LoggedOut() {
        return this.$root.LoggedOut;
      }
      // Getter/setter for the selected profile ID in the root instance
    },
    methods: {
      // async to allow the app to run while waiting for response from axios
      async GetProfiles() { 
        this.$root.GetProfiles();
      },
      ToggleLogIn() {
        this.ShowLogin = !this.ShowLogin;
        this.ShowRegister = !this.ShowRegister;
      },
      // Sends a new profile to the backend to be added to the sql db after validating input
      async AddProfile() {
        if (!this.NewProfile.name || !this.NewProfile.image) {
          alert("Please fill in all fields!");
          return;
        }
        try {
          // request to add profile to the backend
          await axios.post('http://127.0.0.1:8000/api/addprofile', {
            // send a json body to the backend for the request to add profile
            profile_name: this.NewProfile.name,
            profile_image: this.NewProfile.image,
          });
          this.ToggleLogIn();
          // fetch the profiles again from the backend to show the new profile added
        } catch (error) {
          console.error("Failed to add profile:", error);
          alert("Failed to add profile");
        }
        // resets the newprofile values after adding
      },
      async Login() {
        if (!this.LoginId) {
          alert("Please enter an ID!");
          return;
        }
        try {
          const response = await axios.get(`http://127.0.0.1:8000/api/showprofilebyid/${this.LoginId}`);
          this.$root.profile = response.data.data;
          console.log(this.$root.profile);
        } catch(error) {
          console.error("unable to Log in", error);
          alert("unable to Log in");
        }
      },

      logout() {
        this.$root.profile = null;
        this.LoginId = null;
      },

      toggleEditProfile(profile) {
        this.ShowEditProfile = !this.ShowEditProfile;
        this.EditProfile.id = profile.profile_id;
        this.EditProfile.name = profile.profile_name;
        this.EditProfile.image = profile.profile_image;
      },
      async editProfile() {
        if (!this.EditProfile.name || !this.EditProfile.image) {
          alert("Please fill in all fields!");
          return;
        }
        try {
          await axios.post('http://127.0.0.1:8000/api/editprofile', {
            profile_id: this.EditProfile.id,
            profile_name: this.EditProfile.name,
            profile_image: this.EditProfile.image,
          })
        } catch(error) {
          console.error("unable to edit profile", error);
          alert("unable to edit profile");
        }
        this.updateProfile();
        this.cancelEdit();
      },
      cancelEdit() {
        this.EditProfile.id = null;
        this.EditProfile.name = "";
        this.EditProfile.image = "";
        this.ShowEditProfile = false;
      },
      async updateProfile() {
        this.$root.updateProfile();
      }
    },
    
  }
</script>

<style scoped>
  .login-page {
    /* fill the viewport */
    position: fixed;
    z-index: 20;
    inset: 0;
    /* translucent black backdrop */
    background-color: rgba(0, 0, 0, 0.6);
  }

  .profile-container {
    display: flex;
    flex-wrap: wrap;
    gap: 20px;
    padding: 20px;
    justify-content: center;
    background-color: #2c2c2c;
    flex-grow: 1;
  }

  .profile-container .add-profile-button {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 250px;
    height: 150px;
    background-color: #3c3c3c;
    color: white;
    font-size: 50px;
    border-radius: 10px;
    cursor: pointer;
    box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.5);
  }
  .profile-container .add-profile-button:hover {
    background-color: #505050;
  }

  /* Add Profile Form */
  .add-profile-form {
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    display: flex;
    flex-direction: column;
    align-items: left;
    justify-content: center;
    background: #3c3c3c;
    box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.5);
    color: white;
    padding: 20px;
    width: 80%;
    max-width: 400px;
    border-radius: 20px;
  }
  .add-profile-form h3 {
    padding: 10px;
  }
  .add-profile-form input {
    padding: 10px;
    background: black;
    color: white;
    width: 90%;
    border-radius: 10px;
    margin-bottom: 10px;
  }
  .add-profile-form .add-profile-buttons {
    display: flex;
    justify-content: left;
    gap: 10px;
  }
  .add-profile-form .add-profile-buttons button {
    width: 45%;
    height: 40px;
    background-color: #3c3c3c;
    color: white;
    border-radius: 10px;
    cursor: pointer;
    box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.5);
  }
</style>