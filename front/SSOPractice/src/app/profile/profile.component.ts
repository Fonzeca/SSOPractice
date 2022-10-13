import { Component, OnInit } from '@angular/core';
import { interval, Subscription, tap } from 'rxjs';
import { ApiBackService } from '../api-back.service';
import { UserView } from '../model/user-view';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.scss']
})
export class ProfileComponent implements OnInit {

  user: UserView | undefined;


  private subscription: Subscription | undefined;

  public dateNow = new Date();
  public dDay = new Date('Oct 13 2022 00:00:00');
  milliSecondsInASecond = 1000;
  hoursInADay = 24;
  minutesInAnHour = 60;
  SecondsInAMinute = 60;

  public timeDifference: number | undefined;
  public secondsToDday: number | undefined;
  public minutesToDday: number | undefined;
  public hoursToDday: number | undefined;
  public daysToDday: number | undefined;

  constructor(private api: ApiBackService) { }

  ngOnInit(): void {

    this.interact();

    this.subscription = interval(1000)
      .subscribe(x => { this.getTimeDifference(); });
  }

  private getTimeDifference() {
    this.timeDifference = this.dDay.getTime() - new Date().getTime();
    this.allocateTimeUnits(this.timeDifference);
  }
  private allocateTimeUnits(timeDifference: number) {
    this.secondsToDday = Math.floor((timeDifference) / (this.milliSecondsInASecond) % this.SecondsInAMinute);
    this.minutesToDday = Math.floor((timeDifference) / (this.milliSecondsInASecond * this.minutesInAnHour) % this.SecondsInAMinute);
    this.hoursToDday = Math.floor((timeDifference) / (this.milliSecondsInASecond * this.minutesInAnHour * this.SecondsInAMinute));
  }

  interact(): void {
    this.api.getProfile().pipe(
      tap((v) => {
        this.user = v;
        this.dDay = new Date(this.user.expires);
      })
    ).subscribe();
  }

}
