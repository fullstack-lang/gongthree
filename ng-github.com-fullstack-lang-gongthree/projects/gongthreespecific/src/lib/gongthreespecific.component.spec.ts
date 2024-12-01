import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongthreespecificComponent } from './gongthreespecific.component';

describe('GongthreespecificComponent', () => {
  let component: GongthreespecificComponent;
  let fixture: ComponentFixture<GongthreespecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GongthreespecificComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(GongthreespecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
