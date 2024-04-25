<template>
  <v-md-editor :model-value="markdown" mode="preview"></v-md-editor>
</template>

<script>
export default {
  data() {
    return {
      markdown:'![](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/IMG_6522.jpg)\n' +
          '\n' +
          '<font color="#888888" size="3">March 12, 2024</font>\n' +
          '\n' +
          '**<font size="7">月刊（第2期）：眼泪女王</font>**\n' +
          '\n' +
          '本篇是对二〇二四年三月的观剧记录。\n' +
          '\n' +
          '\n' +
          '\n' +
          '#  主要角色介绍：\n' +
          '\n' +
          '\n' +
          '\n' +
          '### 金秀贤（35岁）\n' +
          '\n' +
          '### 皇后集团法律总监/皇后百货法律团队负责人\n' +
          '\n' +
          '![截屏2024-04-04 13.18.44](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/%E6%88%AA%E5%B1%8F2024-04-04%2013.18.44.png)\n' +
          '\n' +
          '毕业于首尔大学法学院，律师。聪明的大脑和视觉效果比财阀更像财阀。当你像模特一样穿着顶级手工鞋、戴着价值一套公寓的手表穿过公司大厅时，你看起来就像刚从曼哈顿上东区的某个地方飞来。他的龙头。当他娶了皇后的女儿时，村里人都像自己的作品一样自豪，说现在是大企业的董事长是国王而不是总统，所以如果像以前一样，我们的贤宇将成为父亲。\n' +
          '\n' +
          '有一次，我嫁给了一个人，因为我相信我爱他们。但三年后，他梦想离婚。绝望地。\n' +
          '\n' +
          '\n' +
          '\n' +
          '### 洪海仁（34岁）  \n' +
          '\n' +
          '### 皇后区集团董事总经理/皇后区百货公司总裁\n' +
          '\n' +
          '![截屏2024-04-04 13.21.30](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/%E6%88%AA%E5%B1%8F2024-04-04%2013.21.30.png)\n' +
          '\n' +
          '一位骄傲的女王，除了蹲着的时候，从来没有弯曲过膝盖。一个仿佛生来就有两张嘴唇，专门用来指挥他人的绝对君主。她是皇后百货商店的老板，这家百货商店像一座城堡一样高高地矗立在江南中部。她的座右铭是“还有什么是做不到的吗？” “我不这么做！”因为你总是这么做。没有太多事情没有成功。目前的目标是进驻百货商场Herkina，加入1万亿俱乐部！\n' +
          '\n' +
          '现在她要做的就是跳到更高的地方，却收到了令人惊讶的消息。\n' +
          '\n' +
          '\n' +
          '\n' +
          '### 尹恩成（34岁）\n' +
          '\n' +
          '### 并购专家\n' +
          '\n' +
          '![截屏2024-04-04 13.22.49](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/%E6%88%AA%E5%B1%8F2024-04-04%2013.22.49.png)\n' +
          '\n' +
          '前华尔街分析师和并购专家。\n' +
          '\n' +
          '在知情人士中，他是一位著名的投资者，他能像算命先生一样预测哪些投资会成功，哪些投资不会成功。这对那些没有镊子导师不知道如何写研究生论文的第三代财阀来说是一个救星。有人千里迢迢飞到纽约来与恩成共进晚餐，也有不少人以自己组建的小团体举办研讨会为借口，花重金邀请恩成。挖掘投资信息。想了解他的人很多，但他的私生活却一无所知，比如他是否结婚，他目前的资产有多少，或者为什么他有时会来去韩国。\n' +
          '\n' +
          '\n' +
          '\n' +
          '### 洪秀哲（33岁）\n' +
          '\n' +
          '### Queens集团执行董事/Queens Mart首席执行官\n' +
          '\n' +
          '![截屏2024-04-04 13.23.55](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/%E6%88%AA%E5%B1%8F2024-04-04%2013.23.55.png)\n' +
          '\n' +
          '海仁的弟弟。厚大虽然很好，但他心里住着一个正在上初二的男孩。从小到大，我几乎是在坚强无比的姐姐的践踏下长大的。虽然我已经长大了，但每当我和姐姐目光接触时，我都会不知不觉地产生警惕，为什么呢？什么！！我无缘无故地生气。我从小就只有一种理想型。这与洪海寅正好相反。这对我来说已经足够了……幸运的是，我找到了这样一个女人，我可以自豪地说，我的婚姻生活比任何人都幸福。\n' +
          '\n' +
          '### 千多惠（38岁）\n' +
          '\n' +
          '### 秀哲的妻子\n' +
          '\n' +
          '![截屏2024-04-04 13.25.16](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/%E6%88%AA%E5%B1%8F2024-04-04%2013.25.16.png)\n' +
          '\n' +
          '优雅奢华的视觉效果，就像刚从美丽岛出来一样。她对秀哲的每一个似乎不像他们说的笑话都有反应并笑，甚至在清晨的早餐时，她第一个出现，打扮得比任何人都完美，让她的爷爷，洪主席高兴极了。他很有礼貌，对时事知识渊博，而且对问题的了解非常渊博，让你想知道一个年轻人怎么会知道这些事情。\n' +
          '\n' +
          '\n' +
          '\n' +
          '# 剧情：\n' +
          '\n' +
          '\n' +
          '\n' +
          '### 第一集\n' +
          '\n' +
          '\n' +
          '\n' +
          '![queenTear1](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/queenTear1.jpg)\n' +
          '\n' +
          '​\t\t  白賢祐和洪海仁的關係陷入緊繃，在職場和家中都摩擦不斷。為了做出攸關自己未來的決定，賢祐回了一趟老家。\t\n' +
          '\n' +
          '\n' +
          '\n' +
          '### 第二集\n' +
          '\n' +
          '\n' +
          '\n' +
          '![queenTear2](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/queenTear2.jpg)\n' +
          '\n' +
          '​\t\t海仁告訴賢祐壞消息之後，賢祐開始對她非常好。另一方面，海仁的老同學回到韓國，進入女王集團。\n' +
          '\n' +
          '\n' +
          '\n' +
          '### 第三集\n' +
          '\n' +
          '\n' +
          '\n' +
          '![queenTear3](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/queenTear3.jpg)\n' +
          '\n' +
          '​    賢祐在海仁被野豬攻擊時救了她，隨後發現令人憂心的證據。與此同時，兩人之間早已熄滅的愛火似乎正慢慢重新燃起。\n' +
          '\n' +
          '\n' +
          '\n' +
          '### 第四集\n' +
          '\n' +
          '\n' +
          '\n' +
          '![queenTear4](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/queenTear4.jpg)\n' +
          '\n' +
          '​\t        \t\t 尹殷盛越來越常介入海仁的生活，這讓賢祐心生嫉妒。隨後，海仁在白斗關需要幫助時出面。\n' +
          '\n' +
          '\n' +
          '\n' +
          '### 第五集\n' +
          '\n' +
          '\n' +
          '\n' +
          '![queenTear5](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/queenTear5.jpg)\n' +
          '\n' +
          '​          \t賢祐和海仁在他的家鄉共度時光，尷尬掩飾逐漸深厚的情感。稍後，海仁獲悉關於自己病情的壞消息。\n' +
          '\n' +
          '\n' +
          '\n' +
          '### 第六集\n' +
          '\n' +
          '\n' +
          '\n' +
          '![queenTear6](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/queenTear6.jpg)\n' +
          '\n' +
          '​\t\t\t                賢祐忽然現身柏林，確保海仁能得到治療。與此同時在韓國，洪滿臺在公司裡面臨危機。\n' +
          '\n' +
          '\n' +
          '\n' +
          '### 第七集\n' +
          '\n' +
          '\n' +
          '\n' +
          '![queenTear7](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/queenTear7.jpg)\n' +
          '\n' +
          '​\t    海仁深感被賢祐背叛後狠狠推開他，態度變得冷冰冰。韓國這一頭，有人策劃對海仁家族不利的陰謀，賢祐抽絲剝繭。\n' +
          '\n' +
          '\n' +
          '\n' +
          '### 第八集\n' +
          '\n' +
          '\n' +
          '\n' +
          '![queenTear8](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/queenTear8.jpg)\n' +
          '\n' +
          '​\t\t\t\t\t     \t縱使與海仁的關係緊張，賢祐依然堅持留在公司，同時也越來越懷疑殷盛。\n' +
          '\n' +
          '\n' +
          '\n' +
          '\n' +
          '\n' +
          '# 直击我心灵的❤️女主\n' +
          '\n' +
          '\n' +
          '\n' +
          '### 拍摄 Y-magazine\n' +
          '\n' +
          '这造型太杀我了\n' +
          '\n' +
          '![](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/IMG_6576.jpg)\n' +
          '\n' +
          '\n' +
          '\n' +
          '\n' +
          '\n' +
          '\n' +
          '\n' +
          '![](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/IMG_6575.jpg)\n' +
          '\n' +
          '\n' +
          '\n' +
          '![](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/IMG_6574.jpg)\n' +
          '\n' +
          '\n' +
          '\n' +
          '![](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/IMG_6587.jpg)\n' +
          '\n' +
          '\n' +
          '\n' +
          '\n' +
          '\n' +
          '### 淡妆版金智媛\n' +
          '\n' +
          '\n' +
          '\n' +
          '![](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/IMG_6578.jpg)\n' +
          '\n' +
          '![](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/IMG_6579.jpg)\n' +
          '\n' +
          '![](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/IMG_6580.jpg)\n' +
          '\n' +
          '\n' +
          '\n' +
          '\n' +
          '\n' +
          '### 可爱版金智媛\n' +
          '\n' +
          '\n' +
          '\n' +
          '![](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/IMG_6582.jpg)\n' +
          '\n' +
          '![](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/IMG_6581.jpg)\n' +
          '\n' +
          '![](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/IMG_6583.jpg)\n' +
          '\n' +
          '![](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/IMG_6585.jpg)\n' +
          '\n' +
          '\n' +
          '\n' +
          '### My favorite\n' +
          '\n' +
          '\n' +
          '\n' +
          '![](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/IMG_6588.jpg)\n' +
          '\n' +
          '\n' +
          '\n' +
          '![](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/%E5%9B%BE%E5%83%8F2024-3-26%2009.08.jpeg)\n' +
          '\n' +
          '\n' +
          '\n' +
          '![](https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/IMG_6424.jpg)'



      // markdown: '# 欢迎使用Markdown编辑器\n' +
      //     '\n' +
      //     '你好！ 这是你第一次使用 **Markdown编辑器** 所展示的欢迎页。如果你想学习如何使用Markdown编辑器, 可以仔细阅读这篇文章，了解一下Markdown的基本语法知识。\n' +
      //     '\n' +
      //     '## 新的改变\n' +
      //     '\n' +
      //     '我们对Markdown编辑器进行了一些功能拓展与语法支持，除了标准的Markdown编辑器功能，我们增加了如下几点新功能，帮助你用它写博客：\n' +
      //     '\n' +
      //     '  1. **全新的界面设计** ，将会带来全新的写作体验；\n' +
      //     '  2. 在创作中心设置你喜爱的代码高亮样式，Markdown **将代码片显示选择的高亮样式** 进行展示；\n' +
      //     '  3. 增加了 **图片拖拽** 功能，你可以将本地的图片直接拖拽到编辑区域直接展示；\n' +
      //     '  4. 全新的 **KaTeX数学公式** 语法；\n' +
      //     '  5. 增加了支持**甘特图的mermaid语法[^1]** 功能；\n' +
      //     '  6. 增加了 **多屏幕编辑** Markdown文章功能；\n' +
      //     '  7. 增加了 **焦点写作模式、预览模式、简洁写作模式、左右区域同步滚轮设置** 等功能，功能按钮位于编辑区域与预览区域中间；\n' +
      //     '  8. 增加了 **检查列表** 功能。\n' +
      //     '\n' +
      //     '[^1]: [mermaid语法说明](https://mermaidjs.github.io/)\n' +
      //     '\n' +
      //     '## 功能快捷键\n' +
      //     '\n' +
      //     '撤销：<kbd>Ctrl/Command</kbd> + <kbd>Z</kbd>\n' +
      //     '重做：<kbd>Ctrl/Command</kbd> + <kbd>Y</kbd>\n' +
      //     '加粗：<kbd>Ctrl/Command</kbd> + <kbd>B</kbd>\n' +
      //     '斜体：<kbd>Ctrl/Command</kbd> + <kbd>I</kbd>\n' +
      //     '标题：<kbd>Ctrl/Command</kbd> + <kbd>Shift</kbd> + <kbd>H</kbd>\n' +
      //     '无序列表：<kbd>Ctrl/Command</kbd> + <kbd>Shift</kbd> + <kbd>U</kbd>\n' +
      //     '有序列表：<kbd>Ctrl/Command</kbd> + <kbd>Shift</kbd> + <kbd>O</kbd>\n' +
      //     '检查列表：<kbd>Ctrl/Command</kbd> + <kbd>Shift</kbd> + <kbd>C</kbd>\n' +
      //     '插入代码：<kbd>Ctrl/Command</kbd> + <kbd>Shift</kbd> + <kbd>K</kbd>\n' +
      //     '插入链接：<kbd>Ctrl/Command</kbd> + <kbd>Shift</kbd> + <kbd>L</kbd>\n' +
      //     '插入图片：<kbd>Ctrl/Command</kbd> + <kbd>Shift</kbd> + <kbd>G</kbd>\n' +
      //     '查找：<kbd>Ctrl/Command</kbd> + <kbd>F</kbd>\n' +
      //     '替换：<kbd>Ctrl/Command</kbd> + <kbd>G</kbd>\n' +
      //     '\n' +
      //     '## 合理的创建标题，有助于目录的生成\n' +
      //     '\n' +
      //     '直接输入1次<kbd>#</kbd>，并按下<kbd>space</kbd>后，将生成1级标题。\n' +
      //     '输入2次<kbd>#</kbd>，并按下<kbd>space</kbd>后，将生成2级标题。\n' +
      //     '以此类推，我们支持6级标题。有助于使用`TOC`语法后生成一个完美的目录。\n' +
      //     '\n' +
      //     '## 如何改变文本的样式\n' +
      //     '\n' +
      //     '*强调文本* _强调文本_\n' +
      //     '\n' +
      //     '**加粗文本** __加粗文本__\n' +
      //     '\n' +
      //     '==标记文本==\n' +
      //     '\n' +
      //     '~~删除文本~~\n' +
      //     '\n' +
      //     '> 引用文本\n' +
      //     '\n' +
      //     'H~2~O is是液体。\n' +
      //     '\n' +
      //     '2^10^ 运算结果是 1024.\n' +
      //     '\n' +
      //     '## 插入链接与图片\n' +
      //     '\n' +
      //     '链接: [link](https://www.csdn.net/).\n' +
      //     '\n' +
      //     '图片: ![Alt](https://imgconvert.csdnimg.cn/aHR0cHM6Ly9hdmF0YXIuY3Nkbi5uZXQvNy83L0IvMV9yYWxmX2h4MTYzY29tLmpwZw)\n' +
      //     '\n' +
      //     '带尺寸的图片: ![Alt](https://imgconvert.csdnimg.cn/aHR0cHM6Ly9hdmF0YXIuY3Nkbi5uZXQvNy83L0IvMV9yYWxmX2h4MTYzY29tLmpwZw =30x30)\n' +
      //     '\n' +
      //     '居中的图片: ![Alt](https://imgconvert.csdnimg.cn/aHR0cHM6Ly9hdmF0YXIuY3Nkbi5uZXQvNy83L0IvMV9yYWxmX2h4MTYzY29tLmpwZw#pic_center)\n' +
      //     '\n' +
      //     '居中并且带尺寸的图片: ![Alt](https://imgconvert.csdnimg.cn/aHR0cHM6Ly9hdmF0YXIuY3Nkbi5uZXQvNy83L0IvMV9yYWxmX2h4MTYzY29tLmpwZw#pic_center =30x30)\n' +
      //     '\n' +
      //     '当然，我们为了让用户更加便捷，我们增加了图片拖拽功能。\n' +
      //     '\n' +
      //     '## 如何插入一段漂亮的代码片\n' +
      //     '\n' +
      //     '去[博客设置](https://mp.csdn.net/console/configBlog)页面，选择一款你喜欢的代码片高亮样式，下面展示同样高亮的 `代码片`.\n' +
      //     '\n' +
      //     '​```javascript\n' +
      //     '// An highlighted block\n' +
      //     'var foo = \'bar\';\n' +
      //     '​```\n' +
      //     '\n' +
      //     '## 生成一个适合你的列表\n' +
      //     '\n' +
      //     '- 项目\n' +
      //     '  - 项目\n' +
      //     '    - 项目\n' +
      //     '\n' +
      //     '1. 项目1\n' +
      //     '2. 项目2\n' +
      //     '3. 项目3\n' +
      //     '\n' +
      //     '- [ ] 计划任务\n' +
      //     '- [x] 完成任务\n' +
      //     '\n' +
      //     '## 创建一个表格\n' +
      //     '\n' +
      //     '一个简单的表格是这么创建的：\n' +
      //     '\n' +
      //     '| 项目 | Value |\n' +
      //     '| ---- | ----- |\n' +
      //     '| 电脑 | $1600 |\n' +
      //     '| 手机 | $12   |\n' +
      //     '| 导管 | $1    |\n' +
      //     '\n' +
      //     '### 设定内容居中、居左、居右\n' +
      //     '\n' +
      //     '使用`:---------:`居中\n' +
      //     '使用`:----------`居左\n' +
      //     '使用`----------:`居右\n' +
      //     '\n' +
      //     '|     第一列     |         第二列 | 第三列         |\n' +
      //     '| :------------: | -------------: | :------------- |\n' +
      //     '| 第一列文本居中 | 第二列文本居右 | 第三列文本居左 |\n' +
      //     '\n' +
      //     '### SmartyPants\n' +
      //     '\n' +
      //     'SmartyPants将ASCII标点字符转换为“智能”印刷标点HTML实体。例如：\n' +
      //     '|    TYPE   |ASCII                          |HTML\n' +
      //     '|----------------|-------------------------------|-----------------------------|\n' +
      //     '|Single backticks|`\'Isn\'t this fun?\'`            |\'Isn\'t this fun?\'            |\n' +
      //     '|Quotes          |`"Isn\'t this fun?"`            |"Isn\'t this fun?"            |\n' +
      //     '|Dashes          |`-- is en-dash, --- is em-dash`|-- is en-dash, --- is em-dash|\n' +
      //     '\n' +
      //     '## 创建一个自定义列表\n' +
      //     '\n' +
      //     'Markdown\n' +
      //     ':  Text-to-HTML conversion tool\n' +
      //     '\n' +
      //     'Authors\n' +
      //     ':  John\n' +
      //     ':  Luke\n' +
      //     '\n' +
      //     '## 如何创建一个注脚\n' +
      //     '\n' +
      //     '一个具有注脚的文本。[^2]\n' +
      //     '\n' +
      //     '[^2]: 注脚的解释\n' +
      //     '\n' +
      //     '##  注释也是必不可少的\n' +
      //     '\n' +
      //     'Markdown将文本转换为 HTML。\n' +
      //     '\n' +
      //     '*[HTML]:   超文本标记语言\n' +
      //     '\n' +
      //     '## KaTeX数学公式\n' +
      //     '\n' +
      //     '您可以使用渲染LaTeX数学表达式 [KaTeX](https://khan.github.io/KaTeX/):\n' +
      //     '\n' +
      //     'Gamma公式展示 $\\Gamma(n) = (n-1)!\\quad\\forall\n' +
      //     'n\\in\\mathbb N$ 是通过欧拉积分\n' +
      //     '\n' +
      //     '$$\n' +
      //     '\\Gamma(z) = \\int_0^\\infty t^{z-1}e^{-t}dt\\,.\n' +
      //     '$$\n' +
      //     '\n' +
      //     '> 你可以找到更多关于的信息 **LaTeX** 数学表达式[here][1].\n' +
      //     '\n' +
      //     '## 新的甘特图功能，丰富你的文章\n' +
      //     '\n' +
      //     '​```mermaid\n' +
      //     'gantt\n' +
      //     '        dateFormat  YYYY-MM-DD\n' +
      //     '        title Adding GANTT diagram functionality to mermaid\n' +
      //     '        section 现有任务\n' +
      //     '        已完成               :done,    des1, 2014-01-06,2014-01-08\n' +
      //     '        进行中               :active,  des2, 2014-01-09, 3d\n' +
      //     '        计划一               :         des3, after des2, 5d\n' +
      //     '        计划二               :         des4, after des3, 5d\n' +
      //     '​```\n' +
      //     '\n' +
      //     '- 关于 **甘特图** 语法，参考 [这儿][2],\n' +
      //     '\n' +
      //     '## UML 图表\n' +
      //     '\n' +
      //     '可以使用UML图表进行渲染。 [Mermaid](https://mermaidjs.github.io/). 例如下面产生的一个序列图：\n' +
      //     '\n' +
      //     '​```mermaid\n' +
      //     'sequenceDiagram\n' +
      //     '张三 ->> 李四: 你好！李四, 最近怎么样?\n' +
      //     '李四-->>王五: 你最近怎么样，王五？\n' +
      //     '李四--x 张三: 我很好，谢谢!\n' +
      //     '李四-x 王五: 我很好，谢谢!\n' +
      //     'Note right of 王五: 李四想了很长时间, 文字太长了<br/>不适合放在一行.\n' +
      //     '\n' +
      //     '李四-->>张三: 打量着王五...\n' +
      //     '张三->>王五: 很好... 王五, 你怎么样?\n' +
      //     '​```\n' +
      //     '\n' +
      //     '这将产生一个流程图。:\n' +
      //     '\n' +
      //     '​```mermaid\n' +
      //     'graph LR\n' +
      //     'A[长方形] -- 链接 --> B((圆))\n' +
      //     'A --> C(圆角长方形)\n' +
      //     'B --> D{菱形}\n' +
      //     'C --> D\n' +
      //     '​```\n' +
      //     '\n' +
      //     '- 关于 **Mermaid** 语法，参考 [这儿][3],\n' +
      //     '\n' +
      //     '## FLowchart流程图\n' +
      //     '\n' +
      //     '我们依旧会支持flowchart的流程图：\n' +
      //     '\n' +
      //     '​```mermaid\n' +
      //     'flowchat\n' +
      //     'st=>start: 开始\n' +
      //     'e=>end: 结束\n' +
      //     'op=>operation: 我的操作\n' +
      //     'cond=>condition: 确认？\n' +
      //     '\n' +
      //     'st->op->cond\n' +
      //     'cond(yes)->e\n' +
      //     'cond(no)->op\n' +
      //     '​```\n' +
      //     '\n' +
      //     '- 关于 **Flowchart流程图** 语法，参考 [这儿][4].\n' +
      //     '\n' +
      //     '## 导出与导入\n' +
      //     '\n' +
      //     '###  导出\n' +
      //     '\n' +
      //     '如果你想尝试使用此编辑器, 你可以在此篇文章任意编辑。当你完成了一篇文章的写作, 在上方工具栏找到 **文章导出** ，生成一个.md文件或者.html文件进行本地保存。\n' +
      //     '\n' +
      //     '### 导入\n' +
      //     '\n' +
      //     '如果你想加载一篇你写过的.md文件，在上方工具栏可以选择导入功能进行对应扩展名的文件导入，\n' +
      //     '继续你的创作。\n' +
      //     '\n' +
      //     '[1]: http://meta.math.stackexchange.com/questions/5020/mathjax-basic-tutorial-and-quick-reference\n' +
      //     '[2]: https://mermaidjs.github.io/\n' +
      //     '[3]: https://mermaidjs.github.io/\n' +
      //     '[4]: http://adrai.github.io/flowchart.js/',
    };
  },
};
</script>

<style>
.article-with-image {
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 400px;
  top:5px;
  background-size: cover;
  background-position: top;
  position: relative;
}
</style>