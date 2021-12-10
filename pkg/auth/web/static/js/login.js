/*
 * @Author: mjzhu
 * @Date: 2021-12-09 10:23:48
 * @LastEditTime: 2021-12-10 17:38:51
 * @FilePath: \static-web\js\login.js
 */
function checkForm() {
    var showLogin = document.getElementById('showLogin')
    var showHome = document.getElementById('showHome')
    var user = document.getElementById('text').value;
    var psw = document.getElementById('password').value;
    var rember = document.getElementById('rember').checked;
    var auto = document.getElementById('auto').checked;
    var error = document.getElementsByClassName('error-msg')[0]
    if (!user) {
        //如果验证不通过
        error.classList.add('showMsg')
        error.innerHTML = '请输入用户名'
        return false;
    } else if (!psw) {
        error.classList.add('showMsg')
        error.innerHTML = '请输入密码'
        return false;
    } else {
        error.classList.remove('showMsg')
        error.innerHTML = ''
        //验证通过
        // self.location='index.html';
        // showLogin.classList.add('hideBlock')
        // showHome.classList.remove('hideBlock')  
        // localStorage.setItem("userName", user);
        return
    }
}

// var login = document.getElementById('login')
// login.addEventListener('click', function () {
//     checkForm()  
// })