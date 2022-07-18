import { Component } from '@angular/core';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent{
  password: string = '';
  username: string = '';
  loginValid: boolean = true;
  constructor() { }

  public onSubmit() {
    if(this.password != "hello"){
      this.loginValid = false;
    }
    console.log(this.password);
    console.log(this.username);
  }
}
