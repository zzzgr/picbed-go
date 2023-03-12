<template>
  <div>
    <n-spin :show="data.loading">
      <n-card :title="data.info.name">
        <!-- uploaders -->
        <n-radio-group v-model:value="data.currUploader">
          <n-space>
            <n-radio v-for="u in data.uploaders" :key="u.id" :value="u.id">
              {{ u.name }}
            </n-radio>
          </n-space>
        </n-radio-group>

        <!-- 上传区 -->
        <n-upload
          class="row-space"
          ref="uploadRef"
          accept="image/jpeg,image/png,image/gif"
          :show-file-list="false"
          :customRequest="customRequest"
        >
          <n-upload-dragger>
            <div style="margin-bottom: 12px">
              <n-icon size="48" :depth="3">
                <CloudUploadOutline />
              </n-icon>
            </div>
            <n-text style="font-size: 16px"> 点击/拖动/粘贴上传 </n-text>
            <n-p depth="3" style="margin: 8px 0 0 0"> 请不要上传敏感图片 </n-p>
          </n-upload-dragger>
        </n-upload>

        <!-- 历史上传 -->
        <div class="row-space" v-show="computedList.length > 0">
          <n-radio-group v-model:value="data.res.mode">
            <n-space>
              <n-radio value="url"> url </n-radio>
              <n-radio value="ubb"> ubb </n-radio>
              <n-radio value="md"> markdown </n-radio>
            </n-space>
          </n-radio-group>

          <div v-for="item in computedList" :key="item.url" class="row-space">
            <div style="font-size: 8px">
              <div>文件: {{ item.filename }}</div>
              <div>大小：{{ item.size }}</div>
              <div>时间：{{ item.time }}</div>
            </div>
            <n-input-group>
              <n-input v-model:value="item.text" size="small" />

              <n-popover trigger="hover">
                <template #trigger>
                  <n-button type="info">预览</n-button>
                </template>
                <n-image
                  :src="item.url"
                  width="128"
                  fallback-src="https://i0.hdslb.com/bfs/album/07310ca14c5a3677d6537136891723a19edf7006.png"
                ></n-image>
              </n-popover>

              <n-button type="primary" @click="copy(item.text)">
                复制
              </n-button>
            </n-input-group>
          </div>
        </div>
      </n-card>
    </n-spin>
  </div>
</template>

<script setup>
import moment from "moment";
import { CloudUploadOutline } from "@vicons/ionicons5";
import { getInfo, uploadPic } from "@/api/open";
import { computed, onMounted, reactive, ref, watch } from "vue";
import { copy, bytesToSize } from "@/util/util";
import { useRoute } from "vue-router/dist/vue-router";

const uploadRef = ref();

const data = reactive({
  info: {
    name: "picbed-go",
  },
  uploaders: [],
  currUploader: "bli",
  loading: false,

  res: {
    list: [
      // { url: "htp:baidu.com", filename: "xx.ing", size: "20", time: "2009" },
      // { url: "http:baidu.com", filename: "xx.ing", size: "20", time: "2009" },
      // { url: "http:baidu.com", filename: "xx.ing", size: "20", time: "2009" },
    ],
    mode: "url",
  },
});

const init = () => {
  getInfo().then((res) => {
    data.info.name = res.data.name;
    data.uploaders = res.data.uploaders;

    let notice = res.data.notice;

    if (notice) {
      window.$notification.info({
        title: "公告",
        content: notice,
        duration: 2500,
        keepAliveOnHover: true,
      });
    }
  });
};

const customRequest = (params) => {
  let formData = new FormData();
  formData.append("file", params.file.file);
  data.loading = true;
  uploadPic(data.currUploader, formData)
    .then((res) => {
      let resData = res.data;
      if (resData) {
        data.res.list.unshift(resData);
      }
      uploadRef.value.clear();
    })
    .finally(() => {
      data.loading = false;
    });
};

const computedList = computed(() => {
  let computedList = data.res.list.map((e) => {
    let { filename, url, size, time } = e;
    let text = url;
    if (data.res.mode == "md") {
      text = `![${filename}](${url})`;
    } else if (data.res.mode == "ubb") {
      text = `[img]${url}[/img]`;
    }

    size = bytesToSize(size);

    time = moment(time * 1000).format("YYYY-MM-DD HH:mm:ss");
    return { filename: filename, url: url, text: text, size: size, time: time };
  });
  return computedList;
});

const route = useRoute();

watch(
  () => route.query,
  (newValue) => {
    let mode = newValue.mode;

    if (["url", "ubb", "md"].includes(mode)) {
      data.res.mode = mode;
    }
  },
  { immediate: true }
);

onMounted(() => {
  init();

  // 监听ctrl+v上传图片
  document.addEventListener("paste", (event) => {
    const items = (event.clipboardData || window.clipboardData).items;
    let file = null;
    if (!items || items.length === 0) {
      window.$message.error("当前浏览器不支持本地");
      return;
    }
    // 搜索剪切板items
    for (let i = 0; i < items.length; i++) {
      if (items[i].type.indexOf("image") !== -1) {
        file = items[i].getAsFile();
        break;
      }
    }
    if (!file) {
      window.$message.warning("粘贴内容非图片");
      return;
    } else {
      customRequest({ file: { file: file } });
    }
  });
});
</script>

<style scoped>
.row-space {
  margin-top: 12px;
}
</style>
