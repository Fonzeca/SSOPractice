import { Component, OnInit } from '@angular/core';
import { faFacebookF, faGooglePlusG } from '@fortawesome/free-brands-svg-icons';
import { map } from 'rxjs';
import { ApiBackService } from '../api-back.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  faFacebookF = faFacebookF
  faGooglePlusG = faGooglePlusG
  
  constructor(private back: ApiBackService) { }

  ngOnInit(): void {
  }

  onSingUpClick() : void{
    document.getElementById('container')?.classList.add("right-panel-active")
  }

  onSingInClick() : void{
    document.getElementById('container')?.classList.remove("right-panel-active")
  }

  onLoginGoogle() : void {
    this.back.googleLogin().pipe(
      map((v) =>{
        window.location.href = v
      })
    )
    .subscribe()
  }

  onFacebookLogin() : void {
    this.back.facebookLogin().pipe(
      map((v) =>{
        window.location.href = v
      })
    )
    .subscribe()
  }

  
}
