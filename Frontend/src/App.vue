<template>
  <div id="app">
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
        <img :src="SelectedProfile.profile_image" alt="Profile Picture">
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

export default {
  name: "App",
  data: function() {
    return {
      // Default profiles array with a single "Guest" user
      profiles: [],
      // Tracks which profile is currently selected
      SelectedProfileID: 1,

      // Sample courses array (can be expanded via Add Course)
      courses: [
        {
          id: 1,
          name: "SIT120 - Responsive Web Apps",
          image: "https://i.imgur.com/4NNunpg.png"
        },
        {
          id: 2,
          name: "SIT111 - Computer Systems",
          image: "https://i.imgur.com/dz3bDfI.png"
        }
      ],
      // Auto-increment ID for newly added courses
      NewCourseID: 3,

      // Sample tasks array with different properties
      tasks: [],
      // Auto-increment ID for newly added tasks
      NewTaskID: 3,

      // Toggles the navbar menu for mobile
      ShowMenu: false
    }
  },
  computed: {
    // Finds the profile object matching the currently selected profile ID
    SelectedProfile() {
      return this.profiles.find(profile => profile.profile_id === this.SelectedProfileID) || {};
    }
  },
  methods: {
    async GetProfiles(){
      try {
        // request profiles from backend using axios and wait for it
        const response = await axios.get('http://127.0.0.1:8000/api/showallprofiles');
        // fill profiles array with data from the db
        this.profiles = response.data.data;
        // if theres an error, alert it
      } catch (error) {
        console.error("Failed to fetch profile:", error);
        alert("failed to fetch profiles");
      }
    },
    async GetCourses() {
      try {
        // get data from backend
        const response = await axios.get('http://127.0.0.1:8000/api/showallcourses');
        // store data in courses array in the component
        this.courses = response.data.data;
      } catch (error) {
        console.error("failed to fetch courses", error);
        alert("failed to fetch courses");
      }
    },
    async GetTasks() {
      try {
        const response = await axios.get('http://127.0.0.1:8000/api/showalltasks')
        this.tasks = response.data.data;
      } catch(error) {
        console.error("failed to fetch tasks", error);
        alert("failed to fetch tasks");
      }
    },
  },
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
