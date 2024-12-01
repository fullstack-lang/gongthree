import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongthreeComponent } from './gongthree.component';

describe('GongthreeComponent', () => {
  let component: GongthreeComponent;
  let fixture: ComponentFixture<GongthreeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GongthreeComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(GongthreeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
