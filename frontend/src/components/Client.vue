<script  setup>
import {reactive} from 'vue'
import { onMounted, onUpdated, onUnmounted } from 'vue'
import {
  Sort,
  Check,
  Delete,
  Edit,
  Message,
  Search,
  Star,
  Download
} from '@element-plus/icons-vue'
import { Timer } from '@element-plus/icons-vue'
// import { isDark } from '~/composables/dark'
// import {Greet} from '../../wailsjs/go/main/App'

onMounted(
  () => {
  window.runtime.EventsOn("msg", (res)=>{
        data.inMsg = res
        data.listFiles = false
  })
  window.runtime.EventsOn("ls", (res)=>{
        var fileNames = res.split(",");
        for(let item of fileNames){
                data.tableData.push(
                  {
                    fileName: item 
                  }
                )
        }
        data.listFiles = true
        data.inMsg = "别看我，看👇!"
  })
  
  }

)

const data = reactive({
  name: "",
  resultText: "Please enter port below 👇",
  port: "0.0.0.0:7000",
  outMsg: "",
  inMsg: "Why not connect?",
  fileNames: "",
  tableData: [{fileName:"1.txt"}],
  listFiles: false
})

function greet() {
    window.go.main.App.Greet(data.name).then(result => {
    data.resultText = result
  })
}

function connect() {
    window.go.main.App.Connect(data.port).then(res => {
         console.log(data.port)
        //  this.port = ""
         data.inMsg = res
       })
}

function send(){
      window.go.main.App.Handle(data.outMsg).then(res => {
        //  this.title = res
        //  data.inMsg = res
          data.outMsg = ""
       })
      // window.runtime.EventsEmit("send", data.outMsg)
 }




const handleRequest = (row) => {
  console.log(row)
  console.log("req " + row.fileName)
  window.go.main.App.Handle("req " + row.fileName)
}

// const tableData: User[] = [
//   {
//     name: 'Tom',
//   },
// ]

</script>

<template>
  <main >
    
      <!-- <div id="result" class="result">
          {{ data.resultText }}
      </div> -->

    <!-- <div id="input1" class="input-box">
      <input id="name" v-model="data.port" autocomplete="off" class="input" type="text"/>
      <button class="btn" @click="connect">Connect</button>
    </div> -->
    <!-- <el-row :gutter="20">
      <el-col :span="12" :offset="3">
              <el-input v-model="input" placeholder="Please input" />

      </el-col>
      <el-col :span="4" :offset="0">
              <el-button :icon="Sort"></el-button>
      </el-col>
    </el-row> -->
    

    <div class="input1">
      <el-row >
        <el-col :span="12" :offset="6">
              <el-input
            v-model="data.port"
            placeholder="输入IP地址:端口">
            <template #append>
              <el-button  :icon="Sort"  @click="connect" />
            </template>
        </el-input>
        </el-col>
      </el-row>
    </div>

    <div id="inMsg" class="result">{{ data.inMsg }}</div>


    <div id="input2" class="msginput">
      <!-- <input id="name" v-model="data.outMsg" autocomplete="on" class="input" type="text" @pressEnter="send"/> -->
      <el-input v-model="data.outMsg" placeholder="Please input" clearable  @keyup.enter.native="send" />
      <!-- <button class="btn" @click="send">Send</button> -->
    </div>

    <div v-if="data.listFiles">
         <el-container>
      <el-main>
         <el-table :data="data.tableData" style="width: 100%">

    <el-table-column label="FileName" width="180">
      <template #default="scope">
           {{ scope.row.fileName }}
      </template>
    </el-table-column>
    <el-table-column label="Operations">
      <template #default="scope">
        <!-- <el-button size="small" @click="handleRequest(scope.row)"
          >Request</el-button
        > -->
        <el-button :icon="Download"  @click="handleRequest(scope.row)" ></el-button>
        <!-- <el-button
          size="small"
          type="danger"
          @click="handleDelete(scope.$index, scope.row)"
          >Delete</el-button
        > -->
      </template>
    </el-table-column>
  </el-table>
      </el-main>
    </el-container>

    </div>
    
  </main>
    <!-- <div>
   <el-row class="mb-4">
       <el-col :span="6" :offset="6">
            <el-input v-model="data.port" placeholder="Please input" />
       </el-col>
       <el-col :span="6" :offset="6">
            <el-button>Default</el-button>
       </el-col>
      
   </el-row>
    </div> -->
</template>

<style scoped >

.input-with-select .el-input-group__prepend {
  background-color: var(--el-fill-color-blank);
}

.main {
  height: 100%;
  width: 100%;
  background: #202124;
  border-radius: 6px;
  border: 0; 
}

.el-button:focus  {
    color: #cfd9df ;
    border-color: #cfd9df;
    background-color: transparent;
    outline: 0;
}
.el-button:hover {
    color: #cfd9df ;
    border-color: #cfd9df;
    background-color: #606d93
}

.el-button {
    --el-button-bg-color: transparent; 
    --el-button-hover-text-color: var(--el-color-primary);
    --el-button-hover-bg-color: var(--el-color-primary-light-9);
    --el-button-hover-border-color: var(--el-color-primary-light-7);
    --el-button-active-text-color: transparent;
    --el-button-active-border-color:  transparent;
    --el-button-active-bg-color: transparent;
}

.el-input__wrapper {
   background-color: transparent;
   
}
.el-main {
  color:  transparent;
}


.el-input__inner {
    color: #e2ebf0;
}
/* .el-input__wrapper.is-focus {
   box-shadow:#d09123
} */
.el-table {
    --el-table-border-color: var(--el-border-color-lighter);
    --el-table-border: 1px solid var(--el-table-border-color);
    --el-table-text-color: var(--el-text-color-regular);
    --el-table-header-text-color: var(--el-text-color-secondary);
    --el-table-row-hover-bg-color: var(--el-fill-color-light);
    --el-table-current-row-bg-color: var(--el-color-primary-light-9);
    --el-table-header-bg-color: #47479400;
    --el-table-fixed-box-shadow: var(--el-box-shadow-light);
    --el-table-bg-color: transparent;
    --el-table-tr-bg-color: transparent;
    --el-table-expanded-cell-bg-color: var(--el-fill-color-blank);
    --el-table-fixed-left-column: inset 10px 0 10px -10px rgba(0, 0, 0, 0.15);
    --el-table-fixed-right-column: inset -10px 0 10px -10px rgba(0, 0, 0, 0.15);
}

/* 这个能生效，.root又不能生效 */
.el-input {
  --el-input-bg-color: transparant;
  --el-input-text-color: #e2ebf0;
  --el-input-border-color:#61718d;
  --el-input-focus-border-color: #e2ebf0;
}


.msginput {
  margin-left: 20px;
  margin-right: 20px;
  
}


.result {
  height: 20px;
  line-height: 20px;
  margin-bottom: 40px;
  margin-top: 120px;
  background: #202124;
  font-size: 30px;
  /* margin: 1.5rem auto; */
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 30px 0 0 20px;
  padding: 0 0px;
  cursor: pointer;
  font-family:'Segoe UI', Tahoma, Geneva, Verdana, sans-serif
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
