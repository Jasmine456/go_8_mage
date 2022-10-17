<script setup>
import {onMounted, reactive, ref} from "vue";
import {DELETE_BLOG, LIST_BLOG,UPDATE_BLOG_STATUS} from "../../api/blog";
import {useRouter} from "vue-router";
import {Message} from "@arco-design/web-vue";


const pagination = reactive({
  total: 50,
  showTotal: true,
  showPageSize: true,
  defaultPageSize: 20,
});

// 当前页发生变化
const handlePageNumberChange = (pn) => {
  console.log(pn)
}

// 分页大小发生变化时
const handlePageSizeChange = (ps) => {
  req.page_size = ps
  req.page_number = 1
  GetDataList();
}

// const loading = ref(false);
const loading = reactive({
  loading: false,
  dot: true,
  tip: "拼命加载中..."
})

const req = reactive({
  keywords:"",
  page_number: 1,
  page_size: pagination.defaultPageSize,
  status:""
})

const blogs = ref([]);

// 拉取数据的函数
const GetDataList = async () => {
  // 开启表格 Loading状态
  loading.loading = true;
  try {
    const resp = await LIST_BLOG(req)
    blogs.value = resp.data.Items;
    pagination.total = resp.data.Total;
    // console.log(resp);
  } catch (error) {
    console.log(error);
  } finally {
    //  结束Loading状态
    loading.loading = false
  }
};
onMounted(async () => {
  GetDataList();
});

const router = useRouter();
//查询详情
const JumpToDetail = (id) => {
  router.push({name: "BlogDetail", params: {id: id}})
}

//删除文章
// 使用bool会导致所有行都处于删除中
//把 deleteLoading设置为对象的id，只有和对象id相等，才为loading
const deleteLoading = ref(-1)

const DeleteBlog = async (id) => {
  deleteLoading.value = id;
  try {
    const resp = await DELETE_BLOG(id)
    // TODO 此次的resp没有返回值，应该为 resp.data.title_name,需要排查此问题
    // console.log("resp内容:"+resp)
    Message.success(`文章：${resp} 已经删除`);
    //删除成功后，需要重新加载数据
    GetDataList();
  } catch (error) {
    console.log(error)
    // Message.error(`删除文章异常：${error.response.data.message}`)
    Message.error(`删除文章异常：${error}`)
  } finally {
    deleteLoading.value = -1;
  }
}

//新增文章
const JumpToCreate =()=>{
  router.push({name:"BlogCreate"});
}

//编辑页面
const JumpToEdit =(id)=>{
  router.push({name:"BlogCreate",query:{id:id}});
};

// 多行选择
const selectedKeys = ref([]);
const rowSelection = reactive({
  type: "checkbox",
  showCheckedAll: true,
  onlyCurrent: false,
});


//表单数据
const form = reactive({
  status:"published"
})
//发布逻辑
const Publish =()=>{
  try{
    console.log(selectedKeys.value)
    const iterator=selectedKeys.value
    for (const id of iterator){
      UPDATE_BLOG_STATUS(id,form)
    }
    GetDataList()
  }catch (error){
    console.log(error)
  }
}

</script>

<template>
  <div class="app">
    <!-- 页头 -->
    <div>
      <a-breadcrumb>
        <a-breadcrumb-item>文章管理</a-breadcrumb-item>
        <a-breadcrumb-item>文章列表</a-breadcrumb-item>
      </a-breadcrumb>
    </div>
    <!-- 表格操作区 -->
    <div class="operate">

      <!--    操作区-->
      <div>
        <a-space>
        <a-button type="primary" size="small" @click="JumpToCreate">
          <template #icon>
            <icon-plus/>
          </template>
          创建
        </a-button>
<!--          <template #cell="{ record }">-->
        <a-button
            size="small"
            type="primary"
            :disabled="selectedKeys.length === 0"
            @click="Publish()"
        >
          <template #icon>
            <icon-send/>
          </template>
          发布
        </a-button>
<!--        </template>-->
        </a-space>


      </div>
      <!--    搜索区-->
      <div class="search">
        <a-input-search
            v-model="req.keywords"
            :style="{width:'320px'}"
            placeholder="请输入文章标题进行搜索"
            @page-change="handlePageNumberChange"
            @page-size-change="handlePageSizeChange"
            search-button
            @search="GetDataList"
            @pressEnter="GetDataList"
        />

      </div>
    </div>

    <!-- 具体的表格-->
    <div class="table">
      <!--        <a-pagination :total="50" show-total show-jumper show-page-size/>-->
      <a-table
          :loading="loading"
          :data="blogs"
          :pagination=pagination
          @page-change="handlePageNumberChange"
          @page-size-change="handlePageSizeChange"
          :row-selection="rowSelection"
          v-model:selected-keys="selectedKeys"
          row-key="id"
      >
        <template #columns>
          <a-table-column title="标题" data-index="title_name"></a-table-column>
          <a-table-column title="作者" data-index="author"></a-table-column>
          <a-table-column title="概要" data-index="summary" ></a-table-column>
          <a-table-column title="状态" data-index="status"></a-table-column>
          <a-table-column title="操作" align="center">
            <template #cell="{ record }">
              <a-space>
                <a-popconfirm
                    :content="`你是否确认要删除文章 ${record.title_name}？`"
                    @ok="DeleteBlog(record.id)"
                >
                <a-button
                    :loading="deleteLoading === record.id"
                    size="small"
                    status="danger">
                  <template #icon>
                    <icon-delete/>
                  </template>
                  删除
                </a-button>
                </a-popconfirm>
                <a-button
                    size="small"
                    type="primary"
                    @click="JumpToEdit(record.id)"
                >
                  <template #icon>
                    <icon-edit/>
                  </template>
                  编辑
                </a-button>

                <a-button
                    size="small"
                    type="primary"
                    @click="JumpToDetail(record.id)"
                >
                  <template #icon>
                    <icon-eye/>
                  </template>
                  预览
                </a-button>
              </a-space>
            </template>
          </a-table-column>
        </template>
      </a-table>
    </div>

  </div>
</template>

<style scoped>
.operate {
  width: 100%;
  margin-top: 12px;
  display: flex;
  align-items: center;
  justify-content: flex-start;
  align-content: center;
}

.search {
  margin-left: auto;
}

.app {
  width: 100%;
}

.table {
  margin-left: 8px;
}

</style>