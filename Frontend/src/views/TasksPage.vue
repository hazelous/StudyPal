<template>
    <!-- 
      Displays tasks with options to filter by course and sort by date/difficulty/weight.
      Includes an "Add Task" modal form and dynamic styling for difficulty, weight, etc.
    -->
  <div>
    <!-- Filter Bar to filter by course and sort by task stats -->
    <div class="filter-bar">
      <div class="dropdown">
        <label for="course-filter">courses:</label>
        <select id="course-filter" v-model="SelectedCourse" @change="UpdateTasks">
          <option value="all">All Courses</option>
          <option
            v-for="course in $root.courses"
            :key="course.course_id"
            :value="course.course_id"
          >
            {{ course.course_name }}
          </option>
        </select>
      </div>
      
      <div class="dropdown">
        <label for="sort-filter">sort:</label>
        <select id="sort-filter" v-model="SortBy" @change="UpdateTasks">
          <option value="due-date">Due Date</option>
          <option value="difficulty">Difficulty</option>
          <option value="weight">Weight</option>
        </select>
      </div>
    </div>
    
    <!-- Task Cards Container -->
    <div class="tasks-container">
      <!-- Display tasks after filtering/sorting -->
      <TaskCard 
      v-for="task in FilteredTasks"
      :key="task.task_id"
      :task="task"
      :course="GetCourseDetails(task.course_id)"
      @toggle-edit="toggleEdit"
      />
      
      <!-- Button to open the Add Task form -->
      <button class="add-task-button" @click="ShowAddTask = true">+</button>

      <!-- Add Task Form -->
      <transition name="slide">
        <div v-if="ShowAddTask" class="add-task-form">
          <h3>Add a New Task</h3>
          <input v-model="NewTask.name" placeholder="Task Name">
          <select v-model="NewTask.CourseID">
            <option disabled value="">Select Course</option>
            <option
              v-for="course in $root.courses"
              :key="course.course_id"
              :value="course.course_id"
            >
              {{ course.course_name }}
            </option>
          </select>
          <input
            v-model.number="NewTask.difficulty"
            type="number"
            placeholder="Difficulty (1-10)"
          >
          <input
            v-model.number="NewTask.weight"
            type="number"
            placeholder="Weight (0-100)"
          >
          <input
            v-model="NewTask.DueDate"
            type="date"
            placeholder="Due Date"
          >
          <div class="task-form-buttons">
            <button @click="AddTask">Add Task</button>
            <button @click="Cancel">Cancel</button>
          </div>
        </div>
      </transition>

      <!-- Edit Task Form -->
      <transition name="slide">
        <div v-if="ShowEdit" class="edit-menu">
          <h3>Edit Course</h3>
          <input v-model="EditTask.name" placeholder="Enter Course Name">
          <select v-model="EditTask.CourseID">
            <option
              v-for="course in $root.courses"
              :key="course.course_id"
              :value="course.course_id"
            >
              {{ course.course_name || "unknown course"}}
            </option>
          </select>
          <input v-model="EditTask.difficulty" type="number" placeholder="Difficulty (1-10)">
          <input v-model="EditTask.weight" type="number" placeholder="Weight (0-100)">
          <input v-model="EditTask.DueDate" type="date" placeholder="Due Date">
          <div class="edit-menu-buttons">
            <button @click.stop="cancelEdit">Cancel</button>
            <button @click.stop="editTask">Edit</button>
          </div>
        </div>
      </transition>

    </div>
  </div>
</template>

<script>
import axios from 'axios'
import TaskCard from '@/components/TaskCard.vue'
export default {
    name: "TasksPage",
    components: { TaskCard },
  data: function() {
    return {
      SelectedCourse: "all", // Filter option (defaults to 'all' courses)
      SortBy: "due-date",    // Sorting criterion (defaults to due-date)
      FilteredTasks: [],     // List of tasks displayed after filter/sort
      ShowAddTask: false,    // Toggles the "Add Task" form
      ShowEdit: false,
      NewTask: {
        name: "",
        CourseID: "",
        difficulty: null,
        weight: null,
        DueDate: "",
        status: "Not Started" // Default status for a new task
      },
      EditTask: {
        id: null,
        name: "",
        CourseID: "",
        difficulty: null,
        weight: null,
        DueDate: "",
      },
    };
  },
  computed: {
    tasks () {
      return this.$root.tasks;
    },
    courses () {
      return this.$root.courses;
    }
  },
  methods: {
    async GetTasks() {
      this.$root.GetTasks();
    },
    async GetCourses() {
      this.$root.GetCourses();
    },
    // Filters tasks by the currently selected course
    FilterTasks() {
      this.FilteredTasks = [];
      if (this.SelectedCourse === "all") {
        this.FilteredTasks = this.tasks;
      } else {
        this.FilteredTasks = this.tasks.filter(task => task.course_id === this.SelectedCourse);
      }
    },
    // Sorts tasks by the selected criterion (due-date, difficulty, weight)
    SortTasks() {
      if (this.SortBy === "due-date") {
        this.FilteredTasks.sort((a, b) => new Date(a.task_due_date).getTime() - new Date(b.task_due_date).getTime());
      } else if (this.SortBy === "difficulty") {
        this.FilteredTasks.sort((a, b) => a.task_difficulty - b.task_difficulty);
      } else if (this.SortBy === "weight") {
        this.FilteredTasks.sort((a, b) => a.task_weight - b.task_weight);
      }
    },
    // Runs both filter and sort whenever a filter/sort option changes
    UpdateTasks() {
      this.FilterTasks();
      this.SortTasks();
    },
    // Fetch details of a course (like name, image) based on its ID
    GetCourseDetails(courseID) {
      return this.courses.find(course => course.course_id === courseID) || {};
    },
    // Returns dynamic styling for difficulty
    GetDifficultyStyle(difficulty) {
      if (difficulty === null) return { backgroundColor: "gray" };
      if (difficulty <= 3) return { backgroundColor: "green" };
      if (difficulty <= 6) return { backgroundColor: "orange", color: "white" };
      return { backgroundColor: "red" };
    },
    // Returns dynamic styling for weight
    GetWeightStyle(weight) {
      if (weight === null) return { backgroundColor: "gray" };
      if (weight < 10) return { backgroundColor: "green" };
      if (weight < 20) return { backgroundColor: "orange", color: "white" };
      return { backgroundColor: "red" };
    },
    // Returns dynamic styling for due date
    GetDueDateStyle(dueDate) {
      const today = new Date();
      dueDate = new Date(dueDate).getTime();
      if (dueDate < today) return { backgroundColor: "red" };
      const oneWeekFromToday = new Date(today);
      oneWeekFromToday.setDate(today.getDate() + 7);
      if (dueDate <= oneWeekFromToday) return { backgroundColor: "orange", color: "white" };
      return { backgroundColor: "gray" };
    },
    // Returns dynamic styling for status
    GetStatusStyle(status) {
      if (status === "Not Started") return { backgroundColor: "gray" };
      if (status === "Working On") return { backgroundColor: "orange", color: "white" };
      if (status === "Completed") return { backgroundColor: "green" };
    },
    // Validates and adds a new task to the global list, then updates the tasks view
    async AddTask() {
      if (!this.NewTask.name) {
        alert("Task name is required!");
        return;
      }
      if (!this.NewTask.CourseID) {
        alert("Please select a course for the task!");
        return;
      }
      if (!this.NewTask.DueDate) {
        alert("Please select a due date for the task!");
        return;
      }

      // Ensure difficulty and weight are within valid ranges or null
      const difficulty =
        this.NewTask.difficulty >= 1 && this.NewTask.difficulty <= 10
          ? this.NewTask.difficulty
          : null;
      const weight =
        this.NewTask.weight >= 0 && this.NewTask.weight <= 100
          ? this.NewTask.weight
          : null;

        // Convert the date-only string into an ISO‐8601 timestamp:
        const isoDate = new Date(this.NewTask.DueDate).toISOString(); 
  // e.g. "2025-04-21T00:00:00.000Z"
      // Add new task to the sql db
      try {
        await axios.post('http://127.0.0.1:8000/api/addtask', {
          task_name: this.NewTask.name,
          course_id: this.NewTask.CourseID,
          task_difficulty: this.NewTask.difficulty,
          task_weight: this.NewTask.weight,
          task_due_date: isoDate,
          task_status: this.NewTask.status,
        })
      } catch(error) {
        console.error("failed to add task", error);
        alert("failed to add task");
      }
      // refresh the tasks page to include new task
      this.GetTasks();
      // Reset form fields and close modal
      this.Cancel();
      // Refresh filter + sort to show the newly added task
      this.UpdateTasks();
    },
    async editTask() {
      const isoDate = new Date(this.EditTask.DueDate).toISOString();
      try {
        await axios.post('http://127.0.0.1:8000/api/edittask', {
          task_id: this.EditTask.id,
          task_name: this.EditTask.name,
          course_id: this.EditTask.CourseID,
          task_difficulty: this.EditTask.difficulty,
          task_weight: this.EditTask.weight,
          task_due_date: isoDate,
        })
      } catch(error) {
        console.error("failed to edit task", error);
        alert("failed to edit task");
      }
      // refresh the tasks page to include changes
      this.GetTasks();
      // Reset form fields and close modal
      this.cancelEdit();
      // Refresh filter + sort to show the new changes
      this.UpdateTasks();
    },
    // Cancels "Add Task," resets fields, and closes the modal
    Cancel() {
      this.NewTask = {
        name: "",
        CourseID: "",
        difficulty: null,
        weight: null,
        DueDate: "",
        status: "Not Started"
      };
      this.ShowAddTask = false;
    },
    cancelEdit() {
      this.ShowEdit = false;
      this.EditTask.id = null;
      this.EditTask.CourseID = null;
      this.EditTask.difficulty = null;
      this.EditTask.weight = null;
      this.EditTask.DueDate = "";
    },
    // Formats date objects to a more readable string format
    FormatDate(raw) {
      // change into a Date object (works if raw is already a Date or an ISO string)
      const d = new Date(raw);

      // If it's not a valid date, return empty string
      if (isNaN(d.getTime())) {
        return "";
      }

      // call toLocaleDateString
      const options = { year: "numeric", month: "short", day: "numeric" };
      return d.toLocaleDateString(undefined, options);
    },
    toggleEdit(task) {
      this.ShowEdit = true;
      this.EditTask.id = task.task_id;
      this.EditTask.name = task.task_name;
      this.EditTask.CourseID = task.course_id;
      this.EditTask.difficulty = task.task_difficulty;
      this.EditTask.weight = task.task_weight;
      // takes only the first 10 characters so only DD-MM-YYYY
      this.EditTask.DueDate = task.task_due_date.substr(0, 10);
    }
  },
  // When the component loads, immediately filter/sort tasks to show them correctly
  async mounted() {
  // wait for root to fetch tasks and courses
  await this.GetTasks();
  await this.GetCourses();
  await this.$root.GetCoursesForSelectedProfile()
  // then apply filter and sorting
  this.UpdateTasks();
},
  watch: {
    // Whenever the root’s tasks array changes, rerun filter+sort
    tasks: 'UpdateTasks',
  },
}
</script>

<style scoped>
/*
  TASKS PAGE
  ----------
  task cards, add-task form, dynamic and reactive styling for difficulty/weight/due date.
*/
.tasks-container {
  display: flex;
  flex-direction: column;
  gap: 15px;
  padding: 20px;
}
/* Add Task Button and Form */
.add-task-button {
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #3c3c3c;
  color: white;
  font-size: 50px;
  border-radius: 10px;
  cursor: pointer;
  box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.5);
  text-align: center;
  font-weight: bold;
}
.add-task-form {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  flex-direction: column;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #3c3c3c;
  box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.5);
  color: white;
  padding: 20px;
  width: 80%;
  max-width: 400px;
  border-radius: 20px;
  z-index: 2;
}
.add-task-form input,
.add-task-form select {
  width: 90%;
  padding: 10px;
  margin: 10px 0;
  background: black;
  color: white;
  border: 1px solid #555;
  border-radius: 5px;
}
.add-task-form .task-form-buttons {
  display: flex;
  justify-content: space-between;
}
.add-task-form .task-form-buttons button {
  width: 45%;
  height: 40px;
  background-color: #3c3c3c;
  color: white;
  border-radius: 10px;
  cursor: pointer;
  box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.5);
}

@media (max-width: 768px) {
  .tasks-container {
    flex-wrap: wrap;
    justify-content: center;
  }
  .task-details,
  .task-stats {
    flex-direction: column;
  }
  .task-title {
    font-size: 16px;
  }
}
</style>