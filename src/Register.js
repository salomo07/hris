function Register() {
	return (
    <div className="row">
		<div className="col s12">
		<div className="container">
		  <div id="register-page" className="row">
		    <div className="col s12 m6 l4 z-depth-4 card-panel border-radius-6 register-card bg-opacity-8">
		      <form className="login-form">
		        <div className="row">
		          <div className="input-field col s12">
		            <h5 className="ml-4">Register</h5>
		            <p className="ml-4">Join to our community now !</p>
		          </div>
		        </div>
		        <div className="row margin">
		          <div className="input-field col s12">
		            <i className="material-icons prefix pt-2">person_outline</i>
		            <input onKeyUp={handleUsername} name="username" id="username" type="text"></input>
		            <label for="username" className="center-align">Username</label>
		          </div>
		        </div>
		        <div className="row margin">
		          <div className="input-field col s12">
		            <i className="material-icons prefix pt-2">mail_outline</i>
		            <input id="email" type="email" onKeyUp={validateEmail}/>
		            <label for="email">Email</label>
		          </div>
		        </div>
		        <div className="row margin">
		          <div className="input-field col s12">
		            <i className="material-icons prefix pt-2">lock_outline</i>
		            <input id="password" type="password"></input>
		            <label for="password">Password</label>
		          </div>
		        </div>
		        <div className="row margin">
		          <div className="input-field col s12">
		            <i className="material-icons prefix pt-2">lock_outline</i>
		            <input id="password-again" type="password"></input>
		            <label for="password-again">Password again</label>
		          </div>
		        </div>
		        <div className="row">
		          <div className="input-field col s12">
		            <a href="index.html" className="btn waves-effect waves-light border-round gradient-45deg-purple-deep-orange col s12">Register</a>
		          </div>
		        </div>
		        <div className="row">
		          <div className="input-field col s12">
		            <p className="margin medium-small">
		              <a href="/login">Already have an account? Login</a>
		            </p>
		          </div>
		        </div>
		      </form>
		    </div>
		  </div>
		</div>
		<div className="content-overlay"></div>
		</div>
	</div>
  );
}
ReactDOM.render(<Register/>,document.getElementById('content'));
$("#username").focus();