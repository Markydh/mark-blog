package utils

import (
	// "context"
	"fmt"
	"mark-blog-backend/global"
	"mark-blog-backend/models"
	"math/rand"
	"time"

	// "math/rand"
	// "time"

	"gopkg.in/gomail.v2"
)



func SendCodeValidate(user models.User)(string,error){
   emailInfo := global.Config.Email
   rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
   checkCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	sendTime := time.Now().Format("2006-01-02 15:04:05")
   	//设置文件发送的内容
      content := fmt.Sprintf(`
      <div style="background:#fff">
          <table width="100%%" border="0" cellspacing="0" cellpadding="0">
              <tbody>
                  <tr style="padding:40px 40px 0 40px;display:table-cell">
                      <td style="font-size:24px;line-height:1.5;color:#000;margin-top:40px">邮箱验证码</td>
                  </tr>
                  <tr>
                      <td style="font-size:14px;color:#333;padding:24px 40px 0 40px">
                          尊敬的 %s 用户您好！<br><br>
                          您于 %s 提交的邮箱验证，<br>
                          您的验证码是：<b> %s </b>,请在 <b> 5 分钟内</b> 进行验证，过期将失效！<br>
                          如果该验证码不是您本人申请，请忽略。
                      </td>
                  </tr>
                  <tr style="padding:40px;display:table-cell"></tr>
              </tbody>
          </table>
      </div>
      <div>
          <table width="100%%" border="0" cellspacing="0" cellpadding="0">
              <tbody>
                  <tr>
                      <td style="padding:20px 40px;font-size:12px;color:#999;line-height:20px;background:#f7f7f7">
                          <a href="{{.SiteAddr}}" style="font-size:14px;color:#929292" rel="noopener" target="_blank">返回 {{.SiteName}} </a>
                      </td>
                  </tr>
              </tbody>
          </table>
      </div>
      `, user.Username, sendTime, checkCode)      

	m := gomail.NewMessage()
	m.SetHeader("From", "2247654142@qq.com")
	m.SetHeader("To", user.Email)
	//抄送人
	//m.SetAddressHeader("Cc", "xxx@qq.com", "xiaozhujiao")
	//主题
	m.SetHeader("Subject", "用户注册验证码")
	//内容
	m.SetBody("text/html", content)

	//拿到token，并进行连接,第4个参数是填授权码
	d := gomail.NewDialer(emailInfo.Host, emailInfo.Port, emailInfo.Account, emailInfo.Password)
	// 发送邮件
	err := d.DialAndSend(m); if err != nil {
		fmt.Printf("DialAndSend err %v:", err)
		panic(err)
	}
	fmt.Printf("send mail success\n")
   return checkCode,err
}
