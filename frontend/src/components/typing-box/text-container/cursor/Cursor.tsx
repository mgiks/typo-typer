export type CursorProps = {
  visible: boolean
  ref: React.Ref<HTMLSpanElement | null>
}

function Cursor({ visible, ref }: CursorProps) {
  const cursor = (
    <span ref={ref} className='text-container__cursor' data-testid='cursor' />
  )

  return visible ? cursor : null
}

export default Cursor
