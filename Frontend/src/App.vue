<template>
  <div id="app">
    <!-- Show Login Page -->
    <LoginPage
      v-if="ShowLogin"
      @login-success="onLoginSuccess"
    />
    <header class="navbar">
      <!-- Logo Section -->
      <div class="logo">
        <img src="https://i.imgur.com/jbC13It.png" alt="StudyPal Logo">
        <div class="title">StudyPal</div>
      </div>

      <!-- Dropdown Menu Toggle (for small screens) -->
      <button class="dropdown-menu-toggle" @click="ShowMenu = !ShowMenu">Menu</button>
      
      <!-- Navigation Links (router-link) -->
      <nav class="nav-links" :class="{ active: ShowMenu }">
        <router-link to="/" active-class="active" exact>Profile</router-link>
        <router-link to="/courses" active-class="active">Courses</router-link>
        <router-link to="/tasks" active-class="active">Tasks</router-link>
        <router-link to="/progress" active-class="active">Progress</router-link>
      </nav>

      <!-- Display Currently Selected User Profile Picture in Navbar -->
      <div class="user">
        <img :src="profile?.profile_image" alt="Profile Picture">
      </div>
    </header>
    
    <!-- Page transition wrapper for router-views -->
    <transition name="page-transition">
      <router-view></router-view>
    </transition>
  </div>
</template>

<script>
import axios from 'axios'
import LoginPage from '@/views/LoginPage.vue'

export default {
  name: "App",
  components: { LoginPage },
  data: function() {
    return {
      // Array to store profile
      profile: {},

      // Array to store courses
      courses: [],
      AllCourses: [],

      // Array to store tasks
      tasks: [],
      TaskStatusList: [],
      // Toggles the navbar menu for mobile
      ShowMenu: false,
      ShowLogin: true
    }
  },

  methods: {
    async GetCourses() {
      try {
        // get data from backend
        const response = await axios.get('http://127.0.0.1:8000/api/showallcourses');
        // store data in courses array in the component
        this.AllCourses = response.data.data;
      } catch (error) {
        console.error("failed to fetch courses", error);
        alert("failed to fetch courses");
      }
    },
    async GetCoursesForSelectedProfile() {
      try {
        const response = await axios.get(`http://127.0.0.1:8000/api/showcourseidforprofile/${this.profile.profile_id}`);
        // ?? [] means "if its null or empty then return an empty array []"
        const courseIDs = response.data.data ?? [];
        this.courses = this.AllCourses.filter(course => courseIDs.includes(course.course_id))
      } catch (error) {
        console.error("failed to fetch courses for profile", error);
        alert("failed to fetch courses for profile");
      }
    },
    async GetTasks() {
      try {
        const response = await axios.get(`http://127.0.0.1:8000/api/showtasksforprofile/${this.profile.profile_id}`)
        this.tasks = response.data.data;
      } catch(error) {
        console.error("failed to fetch tasks", error);
        alert("failed to fetch tasks");
      }
    },
    async GetTaskStatusList() {
      try {
        const response = await axios.get(`http://127.0.0.1:8000/api/showtaskstatusforprofile/${this.profile.profile_id}`)
        this.TaskStatusList = response.data.data;
      } catch(error) {
        console.error("failed to fetch task status'", error);
        alert("failed to fetch task status'");
      }
    },
    async updateProfile() {
        try {
          const response = await axios.get(`http://127.0.0.1:8000/api/showprofilebyid/${this.profile.profile_id}`)
          this.profile = response.data.data
          console.log(this.profile)
        } catch(error) {
          console.error("unable to update profile", error);
          alert("unable to update profile");
        }
    },
  },
  watch: {
    profile() {
      this.GetTaskStatusList();
    }
  },
  mounted() {
    // if you have a default profile selected on startup
    this.GetTaskStatusList();
  }
}
</script>

<style scoped>
/* NAVBAR */
.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: #1c1c1c;
  padding: 5px 10px;
  color: white;
  height: 60px;
}

.navbar .logo {
  display: flex;
  align-items: center;
  gap: 0px;
}

.navbar .logo img {
  width: 40px;
}

.navbar .logo .title {
  font-size: 25px;
  font-weight: bold;
  color: white;
  padding: 5px;
}


/* Navigation Links for large screens */
.nav-links {
  display: flex;
  gap: 20px;
}

/* Dropdown toggle is hidden on large screens */
.dropdown-menu-toggle {
  display: none;
}

.nav-links a {
  margin: 0 15px;
  color: white;
  font-weight: bold;
  text-decoration: none;
  font-size: 15px;
}

.nav-links a.active {
  color: #007BFF; /* highlight active link */
}

.user {
  padding: 5px;
}

.user img {
  width: 40px;
  height: 40px;
  border-radius: 50%;
}

@media (max-width: 768px) {
  .nav-links {
    display: none;
    flex-direction: column;
    background-color: #1c1c1c;
    position: absolute;
    top: 60px;
    left: 0;
    width: 100%;
    padding: 10px 0;
    z-index: 10;
  }
  .nav-links.active {
    display: flex;
  }
  .dropdown-menu-toggle {
    display: flex;
    background-color: #3c3c3c;
    color: white;
    font-size: 5vw;
    border-radius: 10px;
    cursor: pointer;
    box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.5);
    text-align: center;
    font-weight: bold;
    padding: 5px;
    flex-direction: column;
  }
}
</style>
