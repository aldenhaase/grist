import { HttpClient } from '@angular/common/http';
import { Component,} from '@angular/core';
import { Router } from '@angular/router';
import { LoginService } from 'src/app/services/login.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent{
  password: string = '';
  username: string = '';
  loginValid: boolean = false;
  formValid: boolean = false;
  constructor(private http: HttpClient,private router: Router, private loginService: LoginService) { }

  public async onSubmit() {
      this.loginService.sendLoginRequest(this.username, this.password);
   }

public validateForm(){
if(this.password.length < 1 || this.username.length < 1){
      this.formValid = false;
      return
    }
   this.formValid = true; 
}
}
