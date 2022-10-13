import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { tap } from 'rxjs';
import { ApiBackService } from '../api-back.service';

@Component({
  selector: 'app-auth',
  templateUrl: './auth.component.html',
  styleUrls: ['./auth.component.scss']
})
export class AuthComponent implements OnInit {

  constructor(private route: ActivatedRoute, private router: Router, private back: ApiBackService) { }

  ngOnInit(): void {
    this.route.queryParams.subscribe(params => {
      switch (params["type"]) {
        case "google":
          this.back.callBackGoogle(params["code"]).pipe(
            tap((v) => {
              this.router.navigate(['/profile'], { state: { user: v } });
            })
          ).subscribe();
          break
        case "facebook":
            this.back.callBackFacebook(params["code"]).pipe(
              tap((v) => {
                this.router.navigate(['/profile'], { state: { user: v } });
              })
            ).subscribe();
          break
      }
    });
  }

}
