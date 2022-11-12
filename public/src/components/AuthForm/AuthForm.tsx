import { useRef } from "react";

const AuthForm = () => {
  const usernameRef = useRef();
  const passwordRef = useRef();

  return (
    <form>
      <div className="form-group">
        <label htmlFor="username">Username</label>
        <input id="username" type="text" className="form-control" required ref={usernameRef}/>
      </div>
      <div className="form-group">
        <label htmlFor="password">Password</label>
        <input id="password" type="password" className="form-control" required ref={passwordRef}/>
      </div>
      <button type="submit" className="btn btn-success">Submit</button>
    </form>
  );
};

export default AuthForm;