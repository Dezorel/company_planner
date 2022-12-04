/* When the user clicks on the button,
toggle between hiding and showing the dropdown content */
function myFunction() {
  document.getElementById("myDropdown").classList.toggle("show");
}


function filterFunction() {
  var input, filter, ul, li, a, i;
  input = document.getElementById("myInput");
  filter = input.value.toUpperCase();
  div = document.getElementById("myDropdown");
  a = div.getElementsByTagName("a");
  for (i = 0; i < a.length; i++) {
    txtValue = a[i].textContent || a[i].innerText;
    if (txtValue.toUpperCase().indexOf(filter) > -1) {
      a[i].style.display = "";
    } else {
      a[i].style.display = "none";
    }
  }
}
// показать скрыть кабинет
function showCabinet() {
  document.getElementById("calendar__container").setAttribute("style", "display: none");
  document.getElementById("cabinet").setAttribute("style", "display: block");
}
// показать скрыть календарь
function showCalendar() {
  document.getElementById("cabinet").setAttribute("style", "display: none");
  document.getElementById("calendar__container").setAttribute("style", "display: block");
}

// function showPage2() {
//   document.getElementById("page1").setAttribute("style", "display: none");
//   document.getElementById("page2").setAttribute("style", "display: block");
// }

// Show cabinet info
function createCabinet() {
  let cabNum = document.getElementById('cabNum').value
  let cabSize = document.getElementById('cabSize').value
  let optionWiFI
  let optionProector

  return console.log(cabNum)
}

function showPage2(){
  let nameCompany = document.getElementById("company_name").value
  console.log(nameCompany)
  var data = []
      newURL = requestURL+'/company/'+nameCompany
      async function request(){
         data = await sendRequest('GET', newURL)
      }
      request()

      setTimeout(()=>{
         console.log(data)
      },1000)
}
function createCompany(){
  let nameCompany = document.getElementById("company_name").value
  var data = []
      newURL = requestURL+'/company'
      async function request(){
         data = await sendRequest('POST', newURL, {name:nameCompany}) 
      }
      request()

      setTimeout(()=>{
         console.log(data)
      },1000)
}
function addCalendarEvent(){
  let nameCompany = document.getElementById("company_name").value
  let cabNumber = document.getElementById("cabNumber").value
  var data = []
      newURL = requestURL+'/company'
      async function request(){
         data = await sendRequest('POST', newURL, {cabNumber:cabNumber,name:nameCompany})
      }
      request()

      setTimeout(()=>{
         console.log(data)
      },1000)
}



