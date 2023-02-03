import { Component,} from '@angular/core';
import { RegisterService } from 'src/app/services/register.service';
import { UsernameCheckerService } from 'src/app/services/username-checker.service';
import { PasswordCheckerService } from 'src/app/services/password-checker.service';
import { Observable } from 'rxjs';
@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})
export class RegisterComponent{
  public usernameIsValid$ = new Observable<boolean>;
  public passwordIsValid$ = new Observable<boolean>;
  password: string = '';
  username: string = '';
  constructor(private registerService: RegisterService, 
              private usernameService: UsernameCheckerService,
              private passwordService: PasswordCheckerService) { }

  public async onSubmit() {
    this.registerService.sendRegisterRequest(this.username, this.password);
  }

  public checkUsername(){ 
    this.usernameIsValid$ = this.usernameService.checkUsername(this.username);
  }

  public checkPassword(){
    this.passwordIsValid$ = this.passwordService.checkPassword(this.password)
  }


}

