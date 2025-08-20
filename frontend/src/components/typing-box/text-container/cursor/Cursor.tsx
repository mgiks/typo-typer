export type CursorProps = {
  visible: boolean
  ref: React.Ref<HTMLSpanElement | null>
}

function Cursor({ visible, ref }: CursorProps) {
  const cursor = (
    <span
      ref={ref}
      className='text-container__cursor'
      aria-label='Cursor'
    />
  )

  return visible ? cursor : null
}

export default Cursor
