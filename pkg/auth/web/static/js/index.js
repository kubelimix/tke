/*
 * @Author: mjzhu
 * @Date: 2021-12-09 10:23:42
 * @LastEditTime: 2021-12-10 17:33:38
 * @FilePath: \static-web\js\index.js
 */
window.onload = function () {
	toggleStyle()
	Jump()
	// isLogin()
	// loginOut()
}

// 切换样式
function toggleStyle() {
	let filter = document.getElementsByClassName("filter")
	let panelList = document.getElementsByClassName('tab-panel')
	let simpleList = document.getElementsByClassName('c-tab-simple')
	let cTabList = document.getElementsByClassName('c-tab-panel')
	for (let i = 0; i < filter.length; i++) {
		filter[i].addEventListener('click', function () {
			filter[i].classList.add('active')
			panelList[i].classList.add('showBlock')
			for (let j = 0; j < filter.length; j++) {
				if (i != j) {
					filter[j].classList.remove('active')
					panelList[j].classList.remove('showBlock')
				}
			}
		})
	}
	for (let i = 0; i < simpleList.length; i++) {
		simpleList[i].addEventListener('click', function () {
			simpleList[i].classList.add('actived')
			cTabList[i].classList.add('showFlex')
			for (let j = 0; j < simpleList.length; j++) {
				if (i != j) {
					simpleList[j].classList.remove('actived')
					cTabList[j].classList.remove('showFlex')
				}
			}
		})
	}
}

var setNavJump = function (i) {
	let num = 0
	switch (i) {
		case 1:
			num = 710
			break;
		case 3:
			num = 620
			break
		default:
			num = 640
	}
	document.body.scrollTop = i * num;
	document.documentElement.scrollTop = i * num;
	//	双向定位的兼容处理,由于document.body.scrollTop = i*630;可能不生效
}

function Jump() {
	let LiItem = document.getElementsByClassName("li-item")
	for (let i = 0; i < LiItem.length; i++) {
		LiItem[i].addEventListener('click', function () {
			setNavJump(i)
		})
	}
}

// 跳转到登录页面
var showLogin = document.getElementById('showLogin')
var showHome = document.getElementById('showHome')
let jumpBtn = document.getElementById('loginbutton')
jumpBtn.addEventListener('click', function () {
	showLogin.classList.remove('hideBlock')
    showHome.classList.add('hideBlock')
})

// 判断是否登录
function isLogin() {
	var showLogin = document.getElementById('showLogin')
    var showHome = document.getElementById('showHome')
	let loginDom = document.getElementById('login')
	let userInfoDom = document.getElementById('userInfo')
	let userNameDom = document.getElementById('userName')
	let userName = localStorage.getItem('userName')
	if (userName) {
		showLogin.classList.add('hideBlock')
        showHome.classList.remove('hideBlock')
		loginDom.classList.add('disNone')
		userInfoDom.classList.remove('disNone')
		userNameDom.innerHTML = userName
	} else {
		showLogin.classList.remove('hideBlock')
        showHome.classList.add('hideBlock')
		userInfoDom.classList.add('disNone')
		loginDom.classList.remove('disNone')
	}
}

function loginOut() {
	var showLogin = document.getElementById('showLogin')
    var showHome = document.getElementById('showHome')
	let loginOut = document.getElementById('loginOut')
	loginOut.addEventListener('click', function () {
		localStorage.setItem("userName", '');
		// self.location = 'login.html';
		showLogin.classList.remove('hideBlock')
        showHome.classList.add('hideBlock')
	})
}
