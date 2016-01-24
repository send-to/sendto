/* JS for users */
DOM.Ready(function(){
  // User create form fetching of key
  UserCreateForm();

});

function UserCreateForm() {
  DOM.On("#keybase_skip","click",function(e){
    showFields()
    e.preventDefault();
  });
    
  DOM.On("#keybase_button","click",function(e){
    var name = DOM.First('#keybasename').value 
    
    if (name.length > 0) {
      DOM.Get("https://keybase.io/"+name+"/key.asc",function(data) {
        DOM.First('#key').value = data.responseText;
        DOM.First('#name').value = name;
        showFields();
      },function(data){
        DOM.RemoveClass(DOM.First('.warn'),'hidden');
        showFields();
      });
      
    } else {
      alert("Please fill in your keybase.io username before pressing get Key, or press skip to fill in your details manually.")
    }
    
    
    e.preventDefault();
  });

}


function showFields() {
  DOM.RemoveClass(DOM.First('.user_create_fields'),'hidden');
}