<template>
    <!-- 
      Displays overall progress info: 
      - Course progress (percentage of the selected course's tasks marked completed against total tasks)
      - Tasks done
      - Tasks left
      - Tasks overdue
    -->
    <div>
      <!-- Filter bar for selecting a specific course to track progress -->
      <div class="filter-bar">
        <div class="dropdown">
          <label for="course-filter">courses:</label>
          <select id="course-filter" v-model="SelectedCourse">
            <option value="all">All Courses</option>
            <option
              v-for="course in courses"
              :key="course.course_id"
              :value="course.course_id"
            >
              {{ course.course_name }}
            </option>
          </select>
        </div>
      </div>

      <!-- Cards showing progress information -->
      <div class="progress-container">
        <div class="progress-card">
          <div class="progress-title">Course Progress: </div>
          <div class="progress-value course-progress">
            {{ CourseProgress }}%
          </div>
        </div>

        <div class="progress-card">
          <div class="progress-title">Total Tasks Done: </div>
          <div class="progress-value tasks-done">{{ TasksDone }}</div>
        </div>

        <div class="progress-card">
          <div class="progress-title">Total Tasks Left: </div>
          <div class="progress-value tasks-left">{{ TasksLeft }}</div>
        </div>

        <div class="progress-card">
          <div class="progress-title">Tasks Overdue: </div>
          <div class="progress-value tasks-overdue">
            {{ TasksOverdue }}
          </div>
        </div>
      </div>
    </div>
</template>

<script>
export default {
    name: "ProgressPage",
    data: function() {
        return {
        SelectedCourse: "all", // Filter for progress view
        CourseProgress: 0,
        TasksDone: 0,
        TasksLeft: 0,
        TasksOverdue: 0
        };
    },
    computed: {
        // Access the global courses array
        courses() {
        return this.$root.courses;
        },
        // Access the global tasks array
        tasks() {
        return this.$root.tasks;
        },
        // Filter tasks by the selected course
        FilteredTasks() {
        if (this.SelectedCourse === "all") {
            return this.tasks;
        } else {
            return this.tasks.filter(task => task.course_id === this.SelectedCourse);
        }
        }
    },
    methods: {
        // Recompute progress metrics (completed tasks, overdue tasks, etc.)
        UpdateProgress() {
        const totalTasks = this.FilteredTasks.length;
        const completedTasks = this.FilteredTasks.filter(task => task.task_status === "Completed").length;
        const overdueTasks = this.FilteredTasks.filter(task => {
            const today = new Date();
            return new Date(task.task_due_date).getTime() < today && task.task_status !== "Completed";
        }).length;

        // Calculate course progress as a percentage of tasks completed
        this.CourseProgress =
            totalTasks === 0 ? 0 : Math.round((completedTasks / totalTasks) * 100);

        this.TasksDone = completedTasks;
        this.TasksLeft = totalTasks - completedTasks;
        this.TasksOverdue = overdueTasks;
        }
    },
    // Watchers to recalc progress whenever selected course or tasks change
    watch: {
        SelectedCourse() {
            this.UpdateProgress();
        },
        tasks: {
            deep: true, // watch for nested changes in tasks array
            handler() {
                this.UpdateProgress();
            }
        }
    },
    // On mount, calculate initial progress
    mounted() {
        this.$root.GetCourses();
        this.$root.GetTasks();
        this.$root.GetCoursesForSelectedProfile()
        this.UpdateProgress();
    }
}
</script>

<style scoped>
/*
  PROGRESS PAGE
  -------------
  Displays progress cards
*/
.progress-container {
  display: flex;
  flex-wrap: wrap;
  gap: 20px;
  justify-content: flex-start;
  padding: 20px;
  background-color: #2c2c2c;
}
.progress-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 250px;
  height: 200px;
  background-color: #3c3c3c;
  color: white;
  font-size: 18px;
  font-weight: bold;
  border-radius: 10px;
  box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.5);
  cursor: pointer;
}
.progress-title {
  font-size: 20px;
  margin-bottom: 10px;
}
.progress-value {
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 40px;
  font-weight: bold;
  border-radius: 5px;
  padding: 10px 20px;
  margin-top: 10px;
  box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.3);
}
.progress-value.course-progress {
  background-color: green;
}
.progress-value.tasks-done {
  background-color: green;
}
.progress-value.tasks-left {
  background-color: gray;
}
.progress-value.tasks-overdue {
  background-color: red;
}

@media (max-width: 768px) {
  .progress-card {
    width: 90%;
    height: 150px;
  }
}
</style>