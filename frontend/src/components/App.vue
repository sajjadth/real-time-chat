<template>
  <div
    class="container w-100 vh-100 d-flex flex-column justify-content-center align-items-center"
  >
    <div
      class="w-100 vh-100 d-flex flex-column justify-content-center align-items-center"
    >
      <form
        class="w-100 m-4 d-flex flex-column justify-content-center align-items-center"
        @submit="e => handleSubmit(e)"
      >
        <input
          type="text"
          v-model="username"
          class="form-control m-1"
          placeholder="Usernaem"
          autofocus
          required
        />
        <input
          v-model="message"
          type="text"
          class="form-control m-1"
          placeholder="Message"
          required
        />
        <button type="submit" class="btn btn-primary m-1 w-100">
          Send
          <img
            src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABgAAAAYCAYAAADgdz34AAAAAXNSR0IArs4c6QAAAg9JREFUSEuVVYFR4zAQ3HUDpBPylTwdQAmUQCogdAAd5DvgK3jowHQADWiZkxTpZJ8zvGcykSX59m5vtSLaQwDqr24UrwSzeWqct7cfQER7CEItJR+2j23PMvxWYmFt8eQA5l8kvQG4AvAwkS8DYUGKF6mri0MFUvosABaaM4RnUE/kZPOrx1Pexxd6kKQTgd8tUsaBBT9CeOJkQL0uz3Wed1mfd9FvSkr3BB4NIMcuAOfHgE4ADiTnZbSxsWpiKp/XVUl7AP8aQEBLBTbqDuQ0l28vS7wB5MBS7UOP3gsxWbJVJuDETB1ffaLL3gx5KmkP4k7AHYuiHF0VwH9R0F8FHCYD0vp8DBX4RqWkGwODa/zYlhWHJvOjl/jmSV40f0dkIPtdl7BjNV0UefRG8te5vWEFntMkGcBtBTAhVAgvsqGad5J7Wx0qWBYs6QYlsP0vnpGsWsE7mCl6PgtrRVEqUr3NTRZ2YQo1WlZUgf1r9pLV5CPXs1cpKqxHMr3Axx9rKlllGmi1nOQfH7TW2BdADxOnuVhC4BGV0KXZ3aNaRXCIvwCdBJrTzm19dZL/x+wKN18Aj5WKwVVj//HVEM7sgCR9UriqXvBR7DpzvG3XnqI67jfd4kaTkp3EXVYEacG9Cy9Cbd1AkVW4hqxvsSCQz3TbSMuurfWtyyRofq8ssG53H4x8rICDTLZq80l8A26dEi3s/whjAAAAAElFTkSuQmCC"
          />
        </button>
      </form>
      <div
        id="list"
        class="d-flex flex-column w-100 h-100 justify-content-start align-items-start"
      >
        <div v-for="(l, i) in list" :key="i">
          <p id="message" class="w-auto my-1 rounded text-white px-2 py-1">
            <span id="username" class="m-0 text-danger">{{ l.username }}</span>
            <!-- <span class="text-center">.</span> -->
            <span id="date" class="text-white">{{ l.date }}</span
            ><br />
            {{ l.message }}
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
function handleTime() {
  const minute =
    new Date().getMinutes() < 10
      ? "0" + new Date().getMinutes()
      : new Date().getMinutes();
  const hour =
    new Date().getHours() < 10
      ? "0" + new Date().getHours()
      : new Date().getHours();
  return hour + ":" + minute;
}

const socket = new WebSocket("ws://127.0.0.1:8080/ws");
socket.onerror = err => console.log("error:", err);
socket.onclose = close => console.log("close:", close);

export default {
  name: "app",
  data() {
    return {
      list: [],
      username: "",
      message: ""
    };
  },
  methods: {
    handleSubmit: function(e) {
      e.preventDefault();
      const time = handleTime();
      let data = {
        username: this.username,
        message: this.message,
        date: time
      };
      socket.send(JSON.stringify(data));
      this.message = "";
    }
    // handleUsername: function(i, username){
    //   if (i !== 0 || this.list[i].username !== username){
    //     return false
    //   }
    //   return true
    // }
  },
  mounted() {
    socket.onopen = open => console.log("open");
    socket.onmessage = data => this.list.unshift(JSON.parse(data.data));
  }
};
</script>

<!-- CSS libraries -->
<style src="normalize.css/normalize.css"></style>

<!-- Global CSS -->
<style></style>

<!-- Scoped component css -->
<!-- It only affect current component -->
<style scoped>
#user-controller {
  max-height: 30rem;
  max-width: 30rem;
}
#list {
  overflow: auto;
}
#list::-webkit-scrollbar-track {
  -webkit-box-shadow: inset 0 0 1px rgba(0, 0, 0, 0.3);
  background-color: #f5f5f5;
}

#list::-webkit-scrollbar {
  width: 6px;
  background-color: #f5f5f5;
}

#list::-webkit-scrollbar-thumb {
  border-radius: 25px !important;
  background-color: #a7a7a7;
}
#message {
  /* margin-left: 2rem !important; */
  padding-bottom: 1rem !important;
  max-width: 300px;
  overflow-wrap: break-word;
  background: #3f6efc;
  color: #d5fff6;
}

#username {
  margin-right: 0.25rem !important;
  color: rgb(209, 255, 246) !important;
}
#date {
  margin-left: 0.25rem !important;
  color: rgb(209, 255, 246) !important;
  font-size: small;
}
</style>
