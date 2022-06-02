const { defaultTheme } = require("@vuepress/theme-default");

module.exports = {
  lang: "zh-CN",
  title: "你好，微服务",
  description: "一起走进微服务",
  base: "/go-micro/",
  theme: defaultTheme({
    // 在这里进行配置
    sidebar: [
      // SidebarItem
      {
        text: "环境",
        link: "/env.md",
      },
      {
        text: "支付",
        link: "/pay.md",
      },
      {
        text: "开始",
        link: "/",
      },
      {
        text: "缓存",
        link: "/cache.md",
      },
      {
        text: "命令行",
        link: "/console.md",
      },
      {
        text: "熔断",
        link: "/console.md",
      },
      {
        text: "限流",
        link: "/console.md",
      },
    ],
  }),
};
