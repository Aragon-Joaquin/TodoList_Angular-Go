import { Component, ElementRef, OnInit, ViewChild, input } from '@angular/core';

type InputTypes = 'text' | 'password' | 'submit' | 'number';

@Component({
  selector: 'component-input',
  imports: [],
  template: `
    <label for="{{ inputName() }}">
      {{ labelName() }}
      <input
        #inputProps
        type="{{ type() }}"
        id="{{ inputName() }}"
        name="{{ inputName() }}"
      />
    </label>
  `,
  styleUrl: 'input.component.css',
})

/*
! I cant find a way to do proper spreading of attributes.
! I'll be using a ton of inputs at this point. It's frustrating.
*/
export class InputComponent {
  readonly labelName = input<string>('No specified');
  readonly inputName = input.required<string>();
  readonly type = input<InputTypes>('text');
  readonly extraProps = input<Partial<HTMLInputElement>>({});

  // @ViewChild('inputProps', { static: true })
  // inputProps!: ElementRef<HTMLInputElement>;

  // ngOnInit() {
  //   if (this.extraProps == null) return;
  //   const element = this.inputProps?.nativeElement;

  //   Object.keys(this.extraProps).forEach((key) => {
  //     console.log({
  //       '1': element[key as keyof typeof element],
  //       '2': this.extraProps[key as keyof typeof this.extraProps],
  //       element,
  //     });
  //     return (element[key as keyof typeof element] =
  //       this.extraProps[key as keyof typeof this.extraProps]);
  //   });
  //}
}
