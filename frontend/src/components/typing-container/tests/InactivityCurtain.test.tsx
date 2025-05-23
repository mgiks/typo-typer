import '@testing-library/jest-dom'
import { act, render, screen } from '@testing-library/react'
import InactivityCurtain from '../InactivityCurtain'

describe('InactivityCurtain', () => {
  beforeEach(() =>
    render(
      <>
        <InactivityCurtain />
        <textarea data-testid='typing-area' id='typing-area' />{' '}
        <textarea data-testid='some-other-element' id='some-other-area' />
      </>,
    )
  )

  it("should not show when 'typing-area' is focused", () => {
    act(() => {
      screen.getByTestId(/typing-area/).focus()
    })

    const inactivityCurtain = screen.queryByText(
      /click here or type any key to continue/i,
    )

    expect(inactivityCurtain).not.toBeInTheDocument()
  })

  it("should show when 'typing-area' is unfocused", () => {
    act(() => {
      screen.getByTestId(/some-other-element/).focus()
    })

    const inactivityCurtain = screen.queryByText(
      /click here or type any key to continue/i,
    )

    document.getElementById('other-area')?.focus()

    expect(inactivityCurtain).toBeInTheDocument()
  })
})
