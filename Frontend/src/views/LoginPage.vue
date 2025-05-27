<template>
    <div class="login-page">
        <!-- Edit Profile Form -->
        <transition name="slide">
          <div class="edit-menu">
            <h3>Log In</h3>
            <input placeholder="Enter Username">
            <input placeholder="Enter Password">
            <div class="edit-menu-buttons">
                <button>Log In</button>
                <button>Cancel</button>
            </div>
            <a href="#" class="toggle-link">Don't have an Account? Register</a>
          </div>
        </transition>
    </div>
</template>

<script>
    
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
</style>