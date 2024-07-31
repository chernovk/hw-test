package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	currentChannel := in

	for _, stage := range stages {
		temporaryChannel := make(Bi)
		go func(stage Stage, in In, out Bi) {
			defer close(out)
			stageOut := stage(in)
			for {
				select {
				case <-done:
					return
				case item, ok := <-stageOut:
					if !ok {
						return
					}
					select {
					case <-done:
						return
					case out <- item:
					}
				}
			}
		}(stage, currentChannel, temporaryChannel)
		currentChannel = temporaryChannel
	}
	return currentChannel
}
