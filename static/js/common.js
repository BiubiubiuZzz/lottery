$(function(){
	$(".pointer").toggle(
		function(){
			$(this).addClass("pointer2");
			$(".leftTd").hide();
		},
		function(){
			$(this).removeClass("pointer2");
			$(".leftTd").show();
		})
})
function contextPath() {
	var pathName = document.location.pathname;
	var index = pathName.substr(1).indexOf("/");
	var result = pathName.substr(0,index+1);
	return result;
}

/**
 * FrameStaticCheck填充select
 * @param table
 * @param selTag
 * @param selectedValue
 */
function entryOptions(table, selTag, selectedValue,valueName,textName){
	if (typeof selTag == 'string') selTag = $('#' + selTag);
	if(valueName==undefined||textName==undefined){
		for (var i = 0; i < table.length; i++) {
			selTag.addOption(table[i]['checkValue'], table[i]['filedName'], false);
		}
	}else{
		for (var i = 0; i < table.length; i++) {
			selTag.addOption(table[i][valueName], table[i][textName], false);
		}
	}
	if (selectedValue != undefined && selectedValue!='') {
		selTag.selectOptions(selectedValue);
	} else if (selectedValue == '') {
		selTag.addOption('', ' ');
		selTag.selectOptions('');
	}
}
function code2Name(code, table,valueName,textName) {
	if (typeof code === 'undefined') return '';
	if (code === '') return '';

	if (table) {
		if (table instanceof Array) {
			for (var i = 0; i < table.length; i++) {
				if (table[i]['checkValue'] == code) return table[i]['filedName'];
			}
		}
	} else
		return code.toString();
}


function valiSafeStr(str){
	    var filterString = "";  
	    filterString = filterString == "" ? "'~`·!@#$%^&*()-+./" : filterString;  
	    var ch;  
	    var i;  
	    var temp;  
	    var error = false; // 当包含非法字符时，返回True     
	   for (i = 0; i <= (filterString.length - 1); i++) {  
	        ch = filterString.charAt(i);  
	        temp = str.indexOf(ch);  
	        if (temp != -1) {  
	            error = true;  
	            break;  
	        }  
	    }  
	   return error;  
}
function getCheckbox(checkboxName, attrName) {
	var arr = new Array();
	var str = '';
	$('input[name="'+checkboxName+'"]:checked').each(function(){
		str += $(this).attr(attrName)+'#';
	});
	if(str.length>0) {
		str = str.substring(0, str.length-1);
		arr = str.split('#');
	}
	return arr;
}

function getCheckboxStr(checkboxName, attrName) {
	var str = '';
	$('input[name="'+checkboxName+'"]:checked').each(function(){
		str += $(this).attr(attrName)+',';
	});
	if(str.length>0) {
		str = str.substring(0, str.length-1);
	}
	return str;
}

/**
 * 分页控件
 * @param pageCute		当前第几页
 * @param pageTotal		总页数
 * @param dbLine		每页显示数据行
 * @param dbLineCute	当前第几大页
 * @param dbLineTotal	总共多少大页
 * @param dbTotal		总记录数
 * @param funcName		分页控件回调js方法名
 */
function pageWidget(pageCute, pageTotal, dbLine, dbLineCute, dbLineTotal, dbTotal, funcName) {
	var pageStr = '<div>';
	pageStr += '总记录：'+dbTotal+' &nbsp; &nbsp; &nbsp;';
	pageStr += '<a href="javascript:void(0)" onclick="'+funcName+'(1,null,'+dbLine+')">首页</a> &nbsp;'
	if(parseInt(dbLineCute)>1) {
		pageStr += '<a href="javascript:void(0)" onclick="'+funcName+'('+((parseInt(dbLineCute)-1)*10)+',null,'+dbLine+')">&lt;&lt;</a> &nbsp;';
	} else {
		pageStr += '&nbsp;&nbsp;&nbsp;';
	}
	if(parseInt(pageCute)>1) {
		pageStr += '<a href="javascript:void(0)" onclick="'+funcName+'('+(parseInt(pageCute)-1)+',null,'+dbLine+')">&lt;</a> &nbsp;';
	} else {
		pageStr += '&nbsp;&nbsp;&nbsp;';
	}
	for(var i=(dbLineCute-1)*10+1; i<=dbLineCute*10; i++) {
		if(i<=pageTotal) {
			if(i==pageCute) {
				pageStr += '<font color="red">'+i+'</font> &nbsp;';
			} else {
				pageStr += '<a href="javascript:'+funcName+'('+i+',null,'+dbLine+')">'+i+'</a> &nbsp;';
			}
		}
	}
	if(parseInt(pageCute)<pageTotal) {
		pageStr += '<a href="javascript:void(0)" onclick="'+funcName+'('+(parseInt(pageCute)+1)+',null,'+dbLine+')">&gt;</a> &nbsp;';
	} else {
		pageStr += '&nbsp;&nbsp;&nbsp;';
	}
	if(parseInt(dbLineCute)<dbLineTotal) {
		pageStr += '<a href="javascript:void(0)" onclick="'+funcName+'('+(parseInt(dbLineCute)*10+1)+',null,'+dbLine+')">&gt;&gt;</a> &nbsp;';
	} else {
		pageStr += '&nbsp;&nbsp;&nbsp;';
	}
	pageStr += '<a href="javascript:void(0)" onclick="'+funcName+'('+pageTotal+',null,'+dbLine+')">尾页</a> &nbsp;'
		pageStr += '<input type="text" value="'+pageCute+'" id="pageWidgetNum" style="width:23px;height:15px;border:1px solid #C3CED0"/>&nbsp;<input type="button" value=" Go " onclick="'+funcName+'($(\'#pageWidgetNum\').val(),'+pageTotal+','+dbLine+')"/>';
	pageStr += ' &nbsp; 总页数：'+pageTotal;
	pageStr += '<select onchange="'+funcName+'(null,null,this.value)">';
	pageStr += '<option value="5"'+(dbLine==5?' selected':'')+'>5</option>';
	pageStr += '<option value="10"'+(dbLine==10?' selected':'')+'>10</option>';
	pageStr += '<option value="20"'+(dbLine==20?' selected':'')+'>20</option>';
	pageStr += '<option value="50"'+(dbLine==50?' selected':'')+'>50</option>';
	pageStr += '<option value="100"'+(dbLine==100?' selected':'')+'>100</option>';
	pageStr += '<option value="200"'+(dbLine==200?' selected':'')+'>200</option>';
	pageStr += '<option value="500"'+(dbLine==500?' selected':'')+'>500</option>';
	pageStr += '</select>';
	pageStr += '</div>';
	return pageStr;
}

/**
 * 弹出遮罩层
 * @param dialogDivId	需要弹出层的id
 * @param dialogWidth	需要弹出层的宽
 * @param dialogHeight	需要弹出层的高
 * @param title			需要弹出层的标题
 */
function createDialog(dialogDivId, dialogWidth, dialogHeight,title) {
	$('#'+dialogDivId).dialog({
		title:title,
		autoOpen: true,
		width: dialogWidth,
		height: dialogHeight,
		bgiframe: true, //解决ie6中遮罩层盖不住select的问题
		modal:true,//这个就是遮罩效果
		buttons: {
			/*"Ok": function() { 
				$(this).dialog("close"); 
			}, 
			"Cancel": function() { 
				$(this).dialog("close"); 
			}*/
		}
	});
}

/**
 * 页面载入
 * @param objId		需要更新的对象id
 * @param loadUrl	需要加载的url
 * @param vvv		缺省参数
 */
function loadPage(objId, loadUrl, vvv)
{
	if(vvv==null)
	{
		if(loadUrl.indexOf("?")>=0)
		{
			loadUrl += "&loadtime="+Math.random();
		}
		else
		{
			loadUrl += "?loadtime="+Math.random();
		}
	}
	$("#"+objId).load(loadUrl);
}

/**
 * 时间格式化
 * @param char14Str 14位长度的时间字符
 */
function writeDate14(char14Str) {
	return char14Str.substring(0,4)+'-'+char14Str.substring(4,6)+'-'+char14Str.substring(6,8)+' '+char14Str.substring(8,10)+':'+char14Str.substring(10,12)+':'+char14Str.substring(12,14);
}


function changePwd() {
	loadPage('changePwdDiv', contextPath()+'/loadpage/changePwd.jsp');
	createDialog('changePwdDialog', 400, 200, '密码更换');
}

function toChangePwd() {
	if($('#oldPwd').val()=='') {
		alert('请填写原始密码！');
	} else if($('#loginPwd').val()=='') {
		alert('请填写登录密码！');
	} else if($('#loginPwd').val()!=$('#loginPwd2').val()) {
		alert('两次密码输入不一致！');
	} else {
		$.ajax({
			url:contextPath()+'/pwdChange.do',
			type: 'post',
			dataType: 'json',
			cache: true,
			async: true,
			data: {
				"oldPwd":$('#oldPwd').val(),
				"loginPwd":$('#loginPwd').val()
			},
			success:function(res){
				if(res != null)
				{
					alert(res.msg);
					//alert(res.retObj.listObject)
				}
				else{
					alert('请求失败，返回结果null');
				}
			},
			error:function(){
				alert('请求失败 error function');
			}
		});
	}
}

/**
 * 信息提示
 * @param message	提示消息内容
 * @param focusId	焦点id
 */
function alertMessage(message, focusId) {
	alert(message);
	try {$('#'+focusId).focus();} catch(e) {}
}

/**
 * 全选or全取消
 */
function checkOrUnCheck(chkBtn, chkName){
	var flag = false;
	if (chkBtn.checked == true) flag = true;
	$("[name="+chkName+"]:checkbox").attr("checked", flag);
}

/**
 * 单选事件
 */
function radioClick(obj, divObj, value, flag){
	if ((flag==true && obj.value==value) || (flag==false && obj.value!=value))
		$('#'+divObj+'').css("display", "");
	else
		$('#'+divObj+'').css("display", "none");
}

// 日期转换
function formatDate(dateStr){
	return dateStr.substr(0,4)+"."+dateStr.substr(4,2)+"."+dateStr.substr(6,2);
}

/**
 * 比较两个时间的大小
 * 如果开始时间大于结束时间，弹出提示信息，返回false；
 * 否则返回true
 * @param start 开始时间
 * @param end 结束时间
 * @returns {Boolean} 比较结果
 */
function compareTime(start, end){
	  var arrStart = start.split("-");
	  var startDate = new Date(arrStart[0],arrStart[1],arrStart[2]);
	  var startTime = startDate.getTime();
	  
	  var arrEnd = end.split("-");
	  var endDate = new Date(arrEnd[0],arrEnd[1],arrEnd[2]);
	  var endTime = endDate.getTime();
	  
	  if (startTime > endTime){
	    alert("开始时间大于结束时间,请检查.");
		return false;
	  }else{
	    return true;
	  }
	}