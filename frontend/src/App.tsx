import React, { useEffect } from "react";

import { BrowserRouter as Router, Switch, Route } from "react-router-dom";


import Navbar from "./components/Navbar";

import Users from "./components/User";

import Paid from "./components/Paid";

import SignIn from "./components/SignIns";

export default function App() {
  const [token, setToken] = React.useState<String>("");

  useEffect(() => {
    const getToken = localStorage.getItem("uid");
    if (getToken) {
      setToken(getToken);
    }
  }, []);
  console.log("Token",token)

  if (!token) {
    return <SignIn />
  }
 return (

   <Router>

     <div>

       <Navbar />

       <Switch>

         <Route exact path="/" component={Users} />
         <Route exact path="/list" component={Paid} />
        


       </Switch>

     </div>

   </Router>

 );

}