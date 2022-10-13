import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { UserView } from './model/user-view';

@Injectable({
  providedIn: 'root'
})
export class ApiBackService {

  constructor(private http: HttpClient) { }
  
  host = environment.host;


  googleLogin(){
    return this.http.get(this.host + "/google", {responseType: 'text'});
  }

  facebookLogin() {
    return this.http.get(this.host + "/facebook", {responseType: 'text'});
  }

  callBackGoogle(authCode: string){
    return this.http.get(this.host + "/callback?code=" + authCode, {withCredentials:true})
  }

  callBackFacebook(authCode: string){
    return this.http.get(this.host + "/facebook_callback?code=" + authCode, {withCredentials:true})
  }

  getProfile(){
    return this.http.get<UserView>(this.host + "/profile", {withCredentials:true});
  }
  
}
