<template>
    <!-- 
      Displays all courses, plus an "Add Course" button. 
      Shows a modal form to add a new course using transitions.
    -->
    <div class="courses-container">
    <!-- Loop through existing courses and render each one -->
        <CourseCard v-for="course in courses" 
        :key="course.course_id" 
        :course="course" 
        @toggle-edit-course="toggleEdit"/>

        <!-- Button to open the Add Course form -->
        <button class="add-course-button" @click="ShowAddCourse = true">+</button>

        <!-- Add Course Form -->
        <transition name="slide">
            <div v-if="ShowAddCourse" class="add-course-form">
                <h3>Add a New Course</h3>
                <input v-model="NewCourse.name" placeholder="Course Name">
                <input v-model="NewCourse.image" placeholder="Course Image URL">
                <div class="add-course-buttons">
                    <button @click="AddCourse">Add Course</button>
                    <button @click="Cancel">Cancel</button>
                </div>
            </div>
        </transition>
        <!-- Edit Course Form -->
        <transition name="slide">
          <div v-if="ShowEditMenu" class="edit-menu">
            <h3>Edit Course</h3>
            <input v-model="EditCourse.name" placeholder="Enter Course Name">
            <input v-model="EditCourse.image" placeholder="Enter Course Image">
            <div class="edit-menu-buttons">
                <button @click.stop="cancelEdit">Cancel</button>
                <button @click.stop="editCourse">Edit</button>
            </div>
          </div>
        </transition>
    </div>
</template>

<script>
import axios from 'axios'
import CourseCard from '@/components/CourseCard.vue'

    export default {
        name: 'CoursesPage',
        components: { CourseCard },
        data: function() {
            return {
              // Toggles "Add Course" modal
              ShowAddCourse: false,
              ShowEditMenu: false,
              // Stores user input for a new course
              NewCourse: {
                name: "",
                image: ""
              },
              EditCourse: {
                id: null,
                name: "",
                image: ""
              },
            };
        },
        computed: {
            courses() {
            return this.$root.courses;
            }
        },
        methods: {
          async GetCourses() {
            this.$root.GetCourses();
          },
            // Sends a new course to the backend which sends it to sql db, then clears input
          async AddCourse() {
            if (!this.NewCourse.name || !this.NewCourse.image) {
                alert("Please fill in all fields!");
                return;
            }
            try {
                // sends an addcourse post request to backend with the body of the new course
                await axios.post(`http://127.0.0.1:8000/api/addcourse`, {
                course_name: this.NewCourse.name,
                course_image: this.NewCourse.image
                })
            } catch (error) {
                console.error("failed to add course", error);
                alert("failed to add course");
            }
            this.GetCourses();
            this.Cancel();
          },
            // Resets the input fields and closes the modal
          Cancel() {
            this.NewCourse.name = "";
            this.NewCourse.image = "";
            this.ShowAddCourse = false;
          },
          toggleEdit(course) {
            this.ShowEditMenu = true;
            this.EditCourse.id = course.course_id;
            this.EditCourse.name = course.course_name;
            this.EditCourse.image = course.course_image;
          },
          async editCourse() {
            if (!this.EditCourse.name || !this.EditCourse.image) {
              alert("Please fill in all fields!");
              return;
            }
            try {
              await axios.post('http://127.0.0.1:8000/api/editcourse', {
                course_id: this.EditCourse.id,
                course_name: this.EditCourse.name,
                course_image: this.EditCourse.image,
              })
            } catch(error) {
              console.error("unable to edit course", error);
              alert("unable to edit course");
            }
            this.cancelEdit();
            this.GetCourses();
          },
          cancelEdit() {
            this.EditCourse.id = null;
            this.EditCourse.name = "";
            this.EditCourse.image = "";
            this.ShowEditMenu = false;
          }
        },
        mounted () {
            this.GetCourses()
        }
    }
</script>

<style scoped>
    /* 
  COURSES PAGE
  ------------
  Layout for course cards and the add-course form 
*/
.courses-container {
  display: flex;
  flex-wrap: wrap;
  gap: 20px;
  padding: 20px;
  justify-content: center;
  background-color: #2c2c2c;
  flex-grow: 1;
}
.courses-container .add-course-button {
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
.courses-container .add-course-button:hover {
  background-color: #505050;
}
.courses-container .add-course-card {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 250px;
  height: 150px;
  background-color: #3c3c3c;
  color: white;
  font-size: 18px;
  font-weight: bold;
  border-radius: 10px;
  cursor: pointer;
  box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.5);
}
.add-course-form {
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
.add-course-form h3 {
  padding: 10px;
}
.add-course-form input {
  padding: 10px;
  background: black;
  color: white;
  width: 90%;
  border-radius: 10px;
  margin-bottom: 10px;
}
.add-course-form .add-course-buttons {
  display: flex;
  justify-content: left;
  gap: 10px;
}
.add-course-form .add-course-buttons button {
  width: 45%;
  height: 40px;
  background-color: #3c3c3c;
  color: white;
  border-radius: 10px;
  cursor: pointer;
  box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.5);
}

@media (max-width: 768px) {
  .courses-container {
    flex-direction: column;
    align-items: center;
  }
  .courses-container .add-course-button {
    width: 90%;
    height: 150px;
  }
}
</style>