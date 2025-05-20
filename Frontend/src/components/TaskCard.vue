<template>
    <div class="task-card" :style="{ backgroundImage: 'url(' + course.course_image + ')' }">
        <div class="background-overlay"></div>
        <!-- Task Details -->
        <div class="task-details">
          <div class="task-title">{{ task.task_name }}</div>
          <div class="task-course">
            {{ course.course_name || 'Unknown Course' }}
          </div>
        </div>

        <!-- Task Stats (difficulty, weight, due date, status) -->
        <div class="task-stats">
          <div class="stat" :style="GetDifficultyStyle(task.task_difficulty)">
            Difficulty: {{ task.task_difficulty }}
          </div>
          <div class="stat" :style="GetWeightStyle(task.task_weight)">
            Weight: {{ task.task_weight !== null ? task.task_weight + '%' : 'N/A' }}
          </div>
          <div class="stat" :style="GetDueDateStyle(task.task_due_date)">
            Due Date: {{ FormatDate(task.task_due_date) }}
          </div>
          <div class="stat" :style="GetStatusStyle(status)">
            <label :for="'status-' + task.id">Status: </label>
            <select :id="'status-' + task.id" v-model="status" @change="updateStatus">
              <option value="Not Started">Not Started</option>
              <option value="Working On">Working On</option>
              <option value="Completed">Completed</option>
            </select>
          </div>
          <button class="option-menu-button" @click.stop="toggleMenu">☰</button>
        </div>
        <div v-if="showMenu" class="option-menu">
            <ul>
                <li @click.stop="toggleEdit">Edit</li>
                <li @click.stop="deleteTask">Delete</li>
            </ul>
        </div>
      </div>
</template>

<script>
import axios from 'axios'

export default {
    name: 'CourseCard',
    props: {
        task: {
            type: Object,
            required: true
        },
        course: {
            type: Object,
            required: true
        },
        taskStatus: {
            type: string,
            required: true
        }
    },
    data: function() {
        return {
            showMenu: false,
            status: this.taskStatus,
        }
    },
    watch: {
      // if the parent ever changes taskStatus, keep in sync
      taskStatus(newVal) {
        this.localStatus = newVal
      }
    },
    methods: {
        GetDifficultyStyle(difficulty) {
            if (difficulty === null) return { backgroundColor: "gray" };
            if (difficulty <= 3) return { backgroundColor: "green" };
            if (difficulty <= 6) return { backgroundColor: "orange", color: "white" };
            return { backgroundColor: "red" };
        },
        GetWeightStyle(weight) {
            if (weight === null) return { backgroundColor: "gray" };
            if (weight < 10) return { backgroundColor: "green" };
            if (weight < 20) return { backgroundColor: "orange", color: "white" };
            return { backgroundColor: "red" };
        },
        GetDueDateStyle(dueDate) {
            const today = new Date();
            dueDate = new Date(dueDate).getTime();
            if (dueDate < today) return { backgroundColor: "red" };
            const oneWeekFromToday = new Date(today);
            oneWeekFromToday.setDate(today.getDate() + 7);
            if (dueDate <= oneWeekFromToday) return { backgroundColor: "orange", color: "white" };
            return { backgroundColor: "gray" };
        },
        GetStatusStyle(status) {
            if (status === "Not Started") return { backgroundColor: "gray" };
            if (status === "Working On") return { backgroundColor: "orange", color: "white" };
            if (status === "Completed") return { backgroundColor: "green" };
        },
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
        toggleMenu() {
            this.showMenu = !this.showMenu;
        },
        async updateStatus() {
          try {
            await axios.post('http://127.0.0.1:8000/api/updatestatus', {
              "profile_id": this.$root.SelectedProfileID,
              "task_id": this.task.task_id,
              "task_status": this.status,
            })
          } catch (error) {
            console.error("unable to update status", error);
            alert("unable to update status");
          }
        },
        toggleEdit() {
          this.$emit("toggle-edit", this.task);
        },
        async deleteTask() {
          try {
            await axios.post('http://127.0.0.1:8000/api/deletetask', {
              "task_id": this.task.task_id,
            })
          } catch (error) {
            console.error("unable to delete task", error);
            alert("unable to delete task");
          }
          this.$root.GetTasks();
        }
    }
}
</script>

<style scoped>
.task-card {
  position: relative;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 15px;
  background-size: cover;
  background-position: center;
  border-radius: 10px;
  box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.5);
  color: white;
  overflow: visible;
}
.task-card .background-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  border-radius: 10px;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 0;
}
.task-details {
  display: flex;
  flex-direction: column;
  z-index: 1;
}
.task-title {
  font-size: 20px;
  font-weight: bold;
  color: white;
}
.task-course {
  font-size: 12px;
  color: #ccc;
  margin-top: 5px;
}
.task-stats {
  display: flex;
  gap: 10px;
  z-index: 1;
}
.task-status {
  z-index: 1;
}
/* Dynamic styling classes for task stats (difficulty/weight/status/due date) */
.stat {
  padding: 5px 5px;
  font-size: 12px;
  border-radius: 5px;
  min-width: 80px;
  text-align: center;
  font-weight: bold;
  text-shadow: 1px 1px 5px rgba(0, 0, 0, 0.8);
}
.stat select {
  border: none;
  padding: 0px 5px;
  border-radius: 5px;
  font-weight: bold;
  background-color: inherit;
  color: inherit;
  outline: none;
  cursor: pointer;
  text-shadow: 1px 1px 5px rgba(0, 0, 0, 0.8);
}
.stat.difficulty-na,
.stat.weight-na,
.stat.status-not-started,
.stat.due-date-normal {
  background-color: gray;
}
.stat.difficulty-low,
.stat.weight-low,
.stat.status-completed {
  background-color: green;
}
.stat.difficulty-medium,
.stat.weight-medium,
.stat.status-working-on,
.stat.due-date-soon {
  background-color: orange;
  color: white;
}
.stat.difficulty-high,
.stat.weight-high,
.stat.due-date-past {
  background-color: red;
}

/* override the absolute position from the global rule */
.option-menu-button {
  position: static;
}

.option-menu {
  top: 3rem;
  right: 2rem;
  z-index: 2;
}

@media (max-width: 768px) {
  .task-details,
  .task-stats {
    flex-direction: column;
  }
  .task-title {
    font-size: 24px;
  }
  .option-menu-button {
    width: auto;
    height: auto;
    padding: 5px 5px;
    font-size: 12px;
    border-radius: 5px;
    min-width: 80px;
    text-align: center;
    font-weight: bold;
  }
}
</style>