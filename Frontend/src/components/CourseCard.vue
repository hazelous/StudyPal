<template>
    <div class="course-card">
        <img :src="course.course_image" alt="course image">
        <div class="course-name">
            {{ course.course_name }}
        </div>
        <button class="option-menu-button" @click="toggleMenu">☰</button>
        <div v-if="showMenu" class="option-menu">
            <ul>
                <li @click="editCourse">Edit</li>
                <li @click="deleteCourse">Delete</li>
            </ul>
        </div>
    </div>
</template>

<script>
import axios from 'axios'
    export default {
        name: "CourseCard",
        props: {
            course: {
                type: Object,
                required: true,
            }
        },
        data: function() {
            return {
                showMenu: false,
            }
        },
        methods: {
            toggleMenu() {
                this.showMenu = !this.showMenu
            },
            async deleteCourse() {
                try {
                    await axios.post('http://127.0.0.1:8000/api/deletecourse', {
                    course_id: this.course.course_id,
                    })
                    this.$root.GetCourses();
                } catch(error) {
                    console.error('unable to delete course', error);
                    alert('unable to delete course');
                }
            },
            editCourse() {
                this.$emit('toggle-edit-course', this.course);
            }
        }
    }
</script>


<style scoped>
.course-card {
  position: relative;
  width: 250px;
  height: 150px;
  border-radius: 10px;
  cursor: pointer;
  box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.5);
  overflow: hidden;
}
.course-card img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
.course-card .course-name {
  position: absolute;
  width: 100%;
  bottom: 0;
  color: white;
  font-size: 14px;
  font-weight: bold;
  text-shadow: 1px 1px 5px rgba(0, 0, 0, 0.8);
  text-align: center;
  background: rgba(0, 0, 0, 0.5);
  padding: 5px 0;
}

@media (max-width: 768px) {
  .course-card {
    width: 90%;
    height: 150px;
  }
}
</style>