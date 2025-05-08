<!-- src/components/ProfilePage.vue -->
<template>
    <!-- 
      Displays a profile's picture and name. 
      Applies a different styling class if it matches the current active profile.
      Clicking emits 'profile-selected' up to the parent component.
    -->
    <div class="profile-card" 
    :class="{ selected: profile.profile_id === selectedProfileID }"
    @click.stop="selectProfile">
        <img :src="profile.profile_image" alt="Profile Picture" class="profile-picture" />
        <button class="option-menu-button" @click.stop="toggleMenu">☰</button>
        <div v-if="showMenu" class="option-menu">
            <ul>
            <li @click.stop="editProfile">Edit</li>
            <li @click.stop="deleteProfile">Delete</li>
            </ul>
        </div>
        <div class="profile-name">{{ profile.profile_name }}</div>
    </div>
</template>

<script>
import axios from 'axios'

    export default {
        name: 'ProfileCard',
        props: {
            profile: {
                type: Object,
                required: true // A profile object must be passed in
            },
            selectedProfileID: {
                type: Number,
                required: true // The ID of the current selected profile
            }
        },
        data: function() {
            return {
                showMenu: false,
            }
        },
        methods: {
            // Emits an event to the parent, indicating which profile was selected
            selectProfile() {
                this.$emit("profile-selected", this.profile.profile_id);
            },
            // toggles the visibility of the menu
            toggleMenu() {
                this.showMenu = !this.showMenu;
            },
            async deleteProfile() {
                try {
                    await axios.post('http://127.0.0.1:8000/api/deleteprofile', {
                    profile_id: this.profile.profile_id,
                    })
                    this.$root.GetProfiles();
                } catch(error) {
                    console.error('unable to delete profile', error);
                    alert('unable to delete profile');
                }
            },
            editProfile() {
                this.$emit("toggle-edit-profile", this.profile);
            }
        }
    }
</script>

<style scoped>
    .profile-card {
        position: relative;
        width: 250px;
        height: 150px;
        border-radius: 10px;
        cursor: pointer;
        box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.5);
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        overflow: hidden;
        background-color: #3c3c3c;
        color: white;
        text-align: center;
    }

    .profile-card img {
        width: 100%;
        height: 100%;
        object-fit: cover;
        border-radius: 10px;
    }

    .profile-card .profile-name {
        position: absolute;
        bottom: 0;
        width: 100%;
        background: rgba(0, 0, 0, 0.5);
        color: white;
        font-size: 14px;
        font-weight: bold;
        text-align: center;
        padding: 5px 0;
        text-shadow: 1px 1px 5px rgba(0, 0, 0, 0.8);
    }
    .profile-card.selected {
        border: 2px solid blue;
        box-shadow: 0px 4px 15px rgba(0, 0, 255, 0.6);
    }
</style>