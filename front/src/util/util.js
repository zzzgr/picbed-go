// 复制
import useClipboard from "vue-clipboard3";

const { toClipboard } = useClipboard();

export function copy(text) {
  text = String(text);

  if (!text || text == "{}" || text == "[]") {
    return;
  }

  toClipboard(text)
    .then(() => {
      window.$message.success("复制成功");
    })
    .catch((e) => {
      console.error(e);
    });
}

export function bytesToSize(bytes) {
  if (bytes === 0) return "0 B";
  let k = 1024,
    sizes = ["B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"],
    i = Math.floor(Math.log(bytes) / Math.log(k));

  return (bytes / Math.pow(k, i)).toPrecision(3) + " " + sizes[i];
}
