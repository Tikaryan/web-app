$(document).ready(function () {
       var cook = document.cookie.split('|');
       let data = cook[1].split('=')
       let msg = '', errMsg = ''
       if (/msg/.test(data[0])) {
              msg = data[1]
              $('#alert-suc').find('strong').after(msg)
              $('#alert-suc').show();

       }
       else if (/error/.test(data[0])) {
              errMsg = data[1]
              $('#alert-dang').find('strong').after('<p>' + errMsg + '</p>')
              $('#alert-dang').show();
       }
})