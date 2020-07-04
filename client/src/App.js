import React, {Component} from 'react';
import './App.css';
import axios from "axios"

class App extends Component {
  state = {
    email: ""
  }
  
  handleButton = (e) => {
    axios.post("/user", {email: this.state.email})
      .then(
        res => {
          this.getUsers()
        }
      );
  }

  handleInputChange = (e) => {
    this.setState({
      email: e.target.value
    })
  }

  getUsers = () => {
    axios.get("/user")
      .then(
        res => {
          console.log(res.data);
          let items = res.data.map(usr => {
            let source = `data:image/png;base64,${usr.indenticon}`
            return <li key={usr.email} className="listItem">{usr.email}: <div/> <img src={source} alt="what" /></li>
          })
          let images = <ul className="App">{items}</ul>
          this.setState({
            images: images
          })
        }
      );
  }

  componentDidMount() {
    this.getUsers()
  }

  render() {    
    return (
      <div className="App">
        <div className="userForm">
          <label>
            Enter an email:
            <input type="text" value={this.state.email} onChange={this.handleInputChange} />
            <button onClick={this.handleButton}>Create user</button>
          </label>
        </div>
        {this.state.images}
      </div>
    )
  };
}

export default App;
